package writer

import (
	"d2-item-sorter/pkg/data"
	"d2-item-sorter/pkg/domain"
	"d2-item-sorter/pkg/internal/utils"
	"math"
)

func manipulateItemPosition(originalData []byte, x int, y int) []byte {
	originalData = utils.WriteBits(originalData, 65, 4, x)
	originalData = utils.WriteBits(originalData, 69, 4, y)
	return originalData
}

func PrepareDataToWrite(s domain.Stash) ([]byte, int) {

	dataToWrite := []byte{}

	dataToWrite = append(dataToWrite, []byte(s.StashType)...)
	dataToWrite = append(dataToWrite, s.ExtraByte)
	dataToWrite = append(dataToWrite, []byte(s.Version)...)
	dataToWrite = append(dataToWrite, s.SharedGold.Bytes...)
	dataToWrite = append(dataToWrite, byte(len(s.StashPages)), 0x00, 0x00, 0x00)

	itemCount := 0

	for _, page := range s.StashPages {

		//Stash Page beginning
		dataToWrite = append(dataToWrite, []byte("ST")...)
		//plugy stash flags (0x01 = normal, 0x03 = yellow index, 0x07 = red index)
		dataToWrite = append(dataToWrite, data.StashPageIndexMap[s.StashType][page.Type], 0x00, 0x00, 0x00)
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
			dataToWrite = append(dataToWrite, manipulateItemPosition(item.GetOriginalData(), item.StoredAt.X, item.StoredAt.Y)...)
		}
	}
	return dataToWrite, itemCount
}
