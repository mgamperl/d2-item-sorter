package reader

import (
	"bufio"
	"bytes"
	"d2-item-sorter/pkg/data"
	"d2-item-sorter/pkg/domain"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nokka/d2s"
)

type StashNum struct {
	Count uint16
}

func ReloadAndReturnStash(stash domain.Stash) domain.Stash {
	s, _ := readStash(stash.Filepath)
	return *s
}

func ReloadStash(stash *domain.Stash) error {
	stash, err := readStash(stash.Filepath)
	return err
}

func ReadStashFromFile(path string) (*domain.Stash, error) {
	stash, error := readStash(path)
	return stash, error
}

func readStash(path string) (*domain.Stash, error) {
	s := domain.Stash{}
	s.Filepath = path
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error while opening file %s", path)
		return nil, err
	}

	defer file.Close()

	bfr := bufio.NewReader(file)

	s.Bfr = bfr

	s.StashType, s.ExtraByte, err = readSharedStashHeader(bfr)

	if err != nil {
		return nil, err
	}

	s.Version, err = readFileVersion(bfr)

	if err != nil {
		return nil, err
	}

	//fmt.Printf("File version: %s\n", s.version)
	if s.Version == "02" || s.StashType == data.StashTypeCharacter {
		s.SharedGold = domain.SharedGold{}
		s.SharedGold.Amount, s.SharedGold.Bytes, _ = readSharedGoldAmount(bfr)
	} else {
		s.SharedGold.Amount = -1
		s.SharedGold.Bytes = []byte{}
	}
	pageCount, err := readNumberOfPages(bfr)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Stash (%s) with %d pages and %d gold\n", s.StashType, pageCount, s.SharedGold.Amount)

	var stashPages []domain.StashPage = []domain.StashPage{}

	var i uint16 = 0
	for ; i < pageCount; i++ {
		stashPage, err := readPage(bfr, s.StashType)
		stashPage.Nr = int(i)
		if err != nil {
			return nil, err
		}
		stashPages = append(stashPages, *stashPage)
		//allItems = append(allItems, stashPage.Items...)
	}

	s.StashPages = stashPages

	return &s, nil
}

func readSharedStashHeader(bfr *bufio.Reader) (data.StashType, byte, error) {

	stashType, stashTypeBytes, err := readStashType(bfr)

	_ = stashTypeBytes

	extraByte, err := bfr.ReadByte()
	if err != nil {
		panic(err)
	}

	return stashType, extraByte, err
}

func readFileVersion(bfr *bufio.Reader) (string, error) {
	p := make([]byte, 2)
	n, err := io.ReadFull(bfr, p)
	if err != nil {
		panic(err)
	}
	if n != 2 {
		return "", fmt.Errorf("could read 2 bytes for file version, got %d bytes", n)
	}

	return string(p), nil
}

func readSharedGoldAmount(bfr *bufio.Reader) (int, []byte, error) {
	res := []byte{}
	p := make([]byte, 4)
	n, err := io.ReadFull(bfr, p)
	if err != nil {
		panic(err)
	}
	if n != 4 {
		return 0, res, fmt.Errorf("could read 4 bytes for number of stashes, got %d bytes", n)
	}

	a := StashNum{}
	//fmt.Printf("P: %v", p)
	err = binary.Read(bytes.NewBuffer(p), binary.LittleEndian, &a)
	if err != nil {
		return 0, res, err
	}
	return int(a.Count), p, nil
}

func readStashType(bfr *bufio.Reader) (data.StashType, []byte, error) {
	res := []byte{}
	p := make([]byte, 3)
	n, err := io.ReadFull(bfr, p)
	if err != nil {
		panic(err)
	}
	if n != 3 {
		return "", res, fmt.Errorf("could read 3 bytes for type of stash, got %d bytes", n)
	}

	return data.StashType(string(p)), p, nil
}

func readNumberOfPages(bfr *bufio.Reader) (uint16, error) {
	p := make([]byte, 4)
	n, err := io.ReadFull(bfr, p)
	if err != nil {
		panic(err)
	}
	if n != 4 {
		return 0, fmt.Errorf("could read 4 bytes for number of stashes, got %d bytes", n)
	}

	a := StashNum{}
	//fmt.Printf("P: %v", p)
	err = binary.Read(bytes.NewBuffer(p), binary.LittleEndian, &a)
	if err != nil {
		return 0, err
	}

	return a.Count, nil
}

func readPage(bfr *bufio.Reader, stashType data.StashType) (*domain.StashPage, error) {
	stashPage := new(domain.StashPage)
	err := readST(bfr)
	if err != nil {
		return nil, err
	}
	// Stash Page Flags
	pageType, _ := bfr.ReadByte()
	bfr.ReadByte()
	bfr.ReadByte()
	bfr.ReadByte()

	stashPage.Type = data.StashPageByteMap[pageType]

	b, err := bfr.ReadByte()

	var nameBytes []byte = []byte{}

	//we have a name for the stash page set, so the item data will follow afterwards
	if b != 0x00 {
		nameBytes = append(nameBytes, b)
		for b != 0x00 {
			b, err = bfr.ReadByte()
			nameBytes = append(nameBytes, b)
		}

		//remove 0x00 at end of string
		if byte(nameBytes[len(nameBytes)-1]) == 0x00 {
			nameBytes = nameBytes[:len(nameBytes)-1]
		}

		stashPage.Name = strings.TrimSpace(string(nameBytes))

		//stashPage.isIgnoredPage = isIgnoredPage(stashPage.Name)
		//stashPage.isPickupPage = isPickupPage(stashPage.Name)
		//stashPage.isSharedPickupPage = isSharedPickupPage(stashPage.Name)

	}

	char := d2s.Character{}

	err = d2s.ParseItems(bfr, &char)
	if err != nil {
		return nil, err
	}

	for _, item := range char.Items {
		newItem := NewItem(item, stashType)
		stashPage.Items = append(stashPage.Items, *newItem)
	}

	return stashPage, err
}

func readST(bfr *bufio.Reader) error {
	p := make([]byte, 2)
	n, err := io.ReadFull(bfr, p)
	if err != nil {
		return err
	}
	if n != 2 {
		return fmt.Errorf("could read 2 bytes for file version, got %d bytes", n)

	}
	return nil
}
