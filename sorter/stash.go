package sorter

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/nokka/d2s"
)

type StashNum struct {
	Count uint16
}

type StashType string

const (
	StashTypeShared    StashType = "SSS"
	StashTypeCharacter StashType = "CST"
	StashTypeUnknown   StashType = ""
)

type SharedGold struct {
	amount int
	bytes  []byte
}

type Stash struct {
	StashType  StashType `json:"stashType"`
	extraByte  byte
	version    string
	SharedGold SharedGold `json:"sharedGold"`
	pageCount  int
	StashPages []StashPage `json:"stashPages"`
	ItemCount  int         `json:"itemCount"`
	bfr        *bufio.Reader
	Filepath   string `json:"filePath"`
}

func (s *Stash) Reload() error {
	return s.ParseStash(s.Filepath)
}

func (s *Stash) ParseStash(path string) error {
	s.Filepath = path
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error while opening file %s", path)
		return err
	}

	defer file.Close()

	bfr := bufio.NewReader(file)

	s.bfr = bfr

	s.StashType, s.extraByte, err = readSharedStashHeader(bfr)

	if err != nil {
		return err
	}

	s.version, err = readFileVersion(bfr)

	if err != nil {
		return err
	}

	//fmt.Printf("File version: %s\n", s.version)
	if s.version == "02" || s.StashType == StashTypeCharacter {
		s.SharedGold = SharedGold{}
		s.SharedGold.amount, s.SharedGold.bytes, _ = readSharedGoldAmount(bfr)
	} else {
		s.SharedGold.amount = -1
		s.SharedGold.bytes = []byte{}
	}
	pageCount, err := readNumberOfPages(bfr)
	s.pageCount = int(pageCount)

	if err != nil {
		return err
	}

	fmt.Printf("Stash (%s) with %d pages and %d gold\n", s.StashType, pageCount, s.SharedGold.amount)

	var stashPages []StashPage = []StashPage{}

	var i uint16 = 0
	for ; i < pageCount; i++ {
		stashPage, err := readPage(bfr, s.StashType)
		stashPage.Nr = i
		if err != nil {
			return err
		}
		stashPages = append(stashPages, *stashPage)
		//allItems = append(allItems, stashPage.Items...)
	}

	s.StashPages = stashPages
	//s.stashPagesToWrite = &stashPages

	return nil
}

func (s *Stash) GetAllItems() []SortableItem {
	allItems := []SortableItem{}
	for _, stashPage := range s.StashPages {
		allItems = append(allItems, stashPage.Items...)
	}
	return allItems
}

func (s *Stash) GetItemsToWrite() []SortableItem {
	allItems := []SortableItem{}
	for _, stashPage := range s.StashPages {
		allItems = append(allItems, stashPage.Items...)
	}
	return allItems
}

func (s *Stash) getNewItemsComparedTo(previousItems []SortableItem) ([]SortableItem, []SortableItem) {

	newItems := []SortableItem{}
	deletedItems := []SortableItem{}
	currentItems := s.GetAllItems()

	fmt.Printf("comparing %d items from previous run to %d items from current run\n", len(previousItems), len(currentItems))

	for _, oldItem := range previousItems {
		exists := false
		if oldItem.D2sItem.ID > 0 {
			for _, currentItem := range currentItems {
				if currentItem.D2sItem.ID > 0 && oldItem.D2sItem.ID == currentItem.D2sItem.ID {
					//fmt.Printf("found item %s\n", currentItem.getName())
					exists = true
					break
				}
			}

			if !exists {
				deletedItems = append(deletedItems, oldItem)
			}
		}
	}

	for _, currentItem := range currentItems {
		exists := false
		if currentItem.D2sItem.ID > 0 {
			for _, oldItem := range previousItems {
				if oldItem.D2sItem.ID > 0 && oldItem.D2sItem.ID == currentItem.D2sItem.ID {
					//fmt.Printf("found item %s", currentItem.getName())
					exists = true
					break
				}
			}

			if !exists {
				newItems = append(newItems, currentItem)
			}
		}
	}

	return newItems, deletedItems
}

func (s *Stash) SortItems(groups []ItemGroup) {
	itemsToSort := []SortableItem{}
	itemsToIgnoreCount := 0
	newStashPages := []StashPage{}

	//fmt.Printf("grouping %d pages\n", len(s.stashPages))

	for idx, page := range s.StashPages {

		if page.isIgnoredPage {
			//ignore this page (just add to new pages without any changes)
			newStashPages = append(newStashPages, page)
			fmt.Printf("\t page %d with %d items will be ignored\n", idx, len(page.Items))
			itemsToIgnoreCount += len(page.Items)
		} else {

			if page.isPickupPage {
				//this is a page we will pickup items but keep the page itself
				newStashPages = append(newStashPages, StashPage{Name: page.Name, Type: page.Type})
				fmt.Printf("\t page %d with %d items will be used as pickup page\n", idx, len(page.Items))
			}
			//if we get until here it means we take the items of the page and add them to all items (base for sorting)
			itemsToSort = append(itemsToSort, page.Items...)
		}
	}

	fmt.Printf("found %d items (%d to sort and %d to ignore)\n", len(itemsToSort)+itemsToIgnoreCount, len(itemsToSort), itemsToIgnoreCount)
	newStashPages = append(newStashPages, groupItems(groups, itemsToSort)...)

	s.StashPages = newStashPages
}

func (s *Stash) TransferPagesTo(target *Stash) {

	ignoredItems := 0
	_ = ignoredItems
	sourceStashPages := []StashPage{}
	targetStashPages := []StashPage{}

	for idx, page := range s.StashPages {

		if page.isIgnoredPage {
			//ignore this page (just add to new pages without any changes)
			sourceStashPages = append(sourceStashPages, page)
			fmt.Printf("page %d with %d items will be ignored\n", idx, len(page.Items))
			ignoredItems += len(page.Items)
		} else {

			if page.isPickupPage {
				//this is a page we will pickup items but keep the page itself
				sourceStashPages = append(sourceStashPages, StashPage{Name: page.Name, Type: page.Type})
				fmt.Printf("page %d with %d items will be used as pickup page\n", idx, len(page.Items))
			}
			//if we get until here it means we take the whole page and put it into the target
			targetStashPages = append(targetStashPages, page)
		}
	}

	//add new (empty) stash pages to our source
	s.StashPages = sourceStashPages
	//move filled pages to our target
	target.StashPages = append(target.StashPages, targetStashPages...)

	fmt.Printf("%d stash pages kept in source and %d pages for target\n", len(s.StashPages), len(targetStashPages))
}

func (s *Stash) prepareDataToWrite() ([]byte, int) {

	/*if s.stashPagesToWrite == nil {
		s.stashPagesToWrite = &s.stashPages
	}
	*/

	dataToWrite := []byte{}

	dataToWrite = append(dataToWrite, []byte(s.StashType)...)
	dataToWrite = append(dataToWrite, s.extraByte)
	dataToWrite = append(dataToWrite, []byte(s.version)...)
	dataToWrite = append(dataToWrite, s.SharedGold.bytes...)
	dataToWrite = append(dataToWrite, byte(len(s.StashPages)), 0x00, 0x00, 0x00)

	itemCount := 0

	for _, page := range s.StashPages {

		//Stash Page beginning
		dataToWrite = append(dataToWrite, []byte("ST")...)
		//plugy stash flags (0x01 = normal, 0x03 = yellow index, 0x07 = red index)
		dataToWrite = append(dataToWrite, StashPageIndexMap[s.StashType][page.Type], 0x00, 0x00, 0x00)
		//stash page name

		if len(page.Name) > 0 {
			//take first 20 characters
			strLength := len(page.Name)
			length := int(math.Min(float64(strLength), 20))
			dataToWrite = append(dataToWrite, []byte(page.Name[:length])...)
		}
		dataToWrite = append(dataToWrite, 0x00)
		//adding header
		dataToWrite = append(dataToWrite, 0x4A, 0x4D, byte(len(page.Items)), 0x0)

		//fmt.Printf("Stash Page (%s) #%d: %d items\n", page.Name, page_idx, len(page.Items))

		for _, item := range page.Items {
			itemCount++
			dataToWrite = append(dataToWrite, item.D2sItem.OriginalData...)
		}
	}
	return dataToWrite, itemCount
}

func (s *Stash) Save(backupSuffix string) error {

	dataToWrite, itemCount := s.prepareDataToWrite()
	err := backupAndWrite(s.Filepath, dataToWrite, backupSuffix)

	fmt.Printf("wrote %d items to: %s\n", itemCount, s.Filepath)
	return err
}

func isIgnoredPage(name string) bool {
	return strings.HasPrefix(name, "(i)")
}

func isPickupPage(name string) bool {
	return strings.HasPrefix(name, "(p)")
}

func isSharedPickupPage(name string) bool {
	return strings.HasPrefix(name, "(sp)")
}

func readSharedStashHeader(bfr *bufio.Reader) (StashType, byte, error) {

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

func readStashType(bfr *bufio.Reader) (StashType, []byte, error) {
	res := []byte{}
	p := make([]byte, 3)
	n, err := io.ReadFull(bfr, p)
	if err != nil {
		panic(err)
	}
	if n != 3 {
		return "", res, fmt.Errorf("could read 3 bytes for type of stash, got %d bytes", n)
	}

	return StashType(string(p)), p, nil
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

func readPage(bfr *bufio.Reader, stashType StashType) (*StashPage, error) {
	stashPage := new(StashPage)
	err := readST(bfr)
	if err != nil {
		return nil, err
	}
	// Stash Page Flags
	pageType, _ := bfr.ReadByte()
	bfr.ReadByte()
	bfr.ReadByte()
	bfr.ReadByte()

	stashPage.Type = StashPageByteMap[pageType]

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

		stashPage.isIgnoredPage = isIgnoredPage(stashPage.Name)
		stashPage.isPickupPage = isPickupPage(stashPage.Name)
		stashPage.isSharedPickupPage = isSharedPickupPage(stashPage.Name)

	}

	char := d2s.Character{}

	err = d2s.ParseItems(bfr, &char)
	if err != nil {
		return nil, err
	}

	for _, item := range char.Items {
		if additionalProperties, ok := itemPropertyMap[item.Type]; ok == true {
			sortableItem := SortableItem{AdditionalProps: *additionalProperties}
			sortableItem.SetItem(item, stashType, false)

			stashPage.Items = append(stashPage.Items, sortableItem)
		} else {
			err = fmt.Errorf("Can not find item with itemCode \"%s\"", item.Type)
			break
		}
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
