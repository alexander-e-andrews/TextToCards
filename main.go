package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"fmt"
	"log"
	"regexp"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

//Sets the text size for text wrapping
var textWidth = 7
var textHeight = 13

//Details for user to set
var wantDuplicates = false
var grouped = true

var repeatColumn = false

//Creates a test imnage of terrible colors
func testImage() *image.RGBA {
	width := 50
	height := 50

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}
	green := color.RGBA{0, 230, 12, 244}
	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if y%2 == 0 {
				img.Set(x, y, green)
			} else {
				img.Set(x, y, cyan)
			}
		}
	}
	return img
}

func hasRepeatColumn(data [][]string) int {
	fmt.Println(data[1][0])
	if data[1][0] == "repeat" {
		fmt.Println("Have a repeat columns")
		return 0
	}
	return -1
}

//Converts and array of strings to an array of ints
func strColToIntCol(col []string) (nums []int) {
	nums = make([]int, len(col))
	for x, strNumber := range col {
		//If there is an error atyoio responds withb a zero so its all good
		nums[x], _ = strconv.Atoi(strNumber)
	}
	return nums
}

//Extraxcts and removes the repeat column from data column and returns both
func removeRepeatColumn(data [][]string, index int) ([][]string, []int) {
	repeats := make([]int, len(data))

	//This will give each row, need to trim off column peice
	//r is row, index is column number
	for r := range data {
		//If there is an error atyoio responds withb a zero so its all good
		repeats[r], _ = strconv.Atoi(data[r][index])

		data[r] = append(data[r][:index], data[r][index+1:]...)
	}
	return data, repeats
}

func main() {
	//Loads in the csv
	lines := start("Book1.csv")
	var repeats []int
	//Was going to if this out, but need it to remove the repeat columns
	repColumnNum := hasRepeatColumn(lines)
	//If we have a repeats columns, extract it
	if repColumnNum >= 0 {
		repeatColumn = true
		lines, repeats = removeRepeatColumn(lines, repColumnNum)
		fmt.Println(repeats)
	}
	craftCards(lines, repeats)
	return
}

//Figure out where text needs to be wrapped around
func textWrapper(label string, x1, x2 int) (labels []string) {
	labels = make([]string, 0)
	//The number of characters that can be held in a line
	charactersHeld := (x2 - x1) / textWidth

	//Keeping a constant loop until break
	//Can do a for loop and the x:characters held chop off but decided to reslice instead as its just easier
	for true {
		//If the line can hold more characters then the label can
		if charactersHeld > len(label) {
			labels = append(labels, label)
			break
		} else {
			labels = append(labels, label[:charactersHeld])
			label = label[charactersHeld:]
		}
	}

	return labels
}

//The x,y 0 0 is top left. Not text size adjusted
func addLabel(img *image.RGBA, dim [4]int, label string) {
	col := color.RGBA{0, 0, 0, 255}
	currentY := dim[1]

	lines := textWrapper(label, dim[0], dim[2])

	for _, line := range lines {
		//Call text wrapper and get array of strings
		point := fixed.Point26_6{fixed.Int26_6(dim[0] * 64), fixed.Int26_6((currentY + textHeight) * 64)}

		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(col),
			Face: basicfont.Face7x13,
			Dot:  point,
		}
		d.DrawString(line)
		currentY += textHeight
	}
}

//Info is an array of arrays, First two arrays will be descriptors
//first array is a location descritot (x1,y1):(x2,y2)
func craftCards(info [][]string, repeats []int) {
	locations := info[0]

	//(x1,y1)(x2,y2)
	points := make([][4]int, 0)
	re := regexp.MustCompile(`(\d+)`)

	fmt.Println(locations)

	//Extract the locations for each text field
	for x, str := range locations {
		//Find 4 numbers from the string
		numbers := re.FindAllString(str, -1)
		fmt.Println(len(numbers))
		if len(numbers) == 4 {
			var loc [4]int
			loc[0], _ = strconv.Atoi(numbers[0])
			loc[1], _ = strconv.Atoi(numbers[1])
			loc[2], _ = strconv.Atoi(numbers[2])
			loc[3], _ = strconv.Atoi(numbers[3])

			points = append(points, loc)
		} else {
			log.Panicf("Wrong number of positional data at index: %d", x)
		}
	}
	//Have the set of destination numbers to place the strings

	tmpImage := testImage()

	//Figure out how many long and bottom to make it
	//numCards = len(info) - 2

	//Currently have 8 cards, ging to do 4*2

	//upLeft := image.Point{0, 0}
	/* iamgeRect := tmpImage.Bounds()
	widt := iamgeRect.Dx() * 4
	hight := iamgeRect.Dy() * 2 */

	//lowRight := image.Point{widt, hight}

	//gridImage := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	//Skip the first 2 lines. First is positiong, second is the titles of each section

	//Cut off the first two rows, as they are header info
	cardTextData := info[2:]
	repeats = repeats[2:]
	//Lets batch into groups of 69, as thats the max number of cards in a group
	cardsToBeSaved := make([]*image.RGBA, 0)
	//number of sheets we have written so far
	writtenGroups := 0

	fmt.Println("Cheese 1")

	for x, y := 0, 0; y < len(cardTextData); {
		fmt.Println("Cheese 2")
		//x should be the count of cards we have
		//y will be the index within the cards
		numDuplicates := 0
		if wantDuplicates {
			numDuplicates = repeats[y]
		} else {
			if repeats[y] > 0 {
				numDuplicates = 1
			}
		}

		if numDuplicates+x > 69 {
			fmt.Println("Cheese 3")
			//We have more cards being added then we can fit
			//add the next couple, then write to a file
			spaceLeft := 69 - x
			//Just add all of them to our card set
			card := writeCard(cardTextData[y], points, tmpImage)
			for t := 0; t < spaceLeft; t++ {
				cardsToBeSaved = append(cardsToBeSaved, card)
			}
			//Dont add to the y
			//y ++
			//Rewrite number to be added
			repeats[y] = repeats[y] - spaceLeft
			x += spaceLeft
			fmt.Println(cardsToBeSaved)
			writeGroupedOrSheet(cardsToBeSaved, writtenGroups)
			writtenGroups++
			cardsToBeSaved = make([]*image.RGBA, 0)
			x = 0
		} else {
			//fmt.Println("Cheese 4")
			//Just add all of them to our card set
			card := writeCard(cardTextData[y], points, tmpImage)
			for t := 0; t < numDuplicates; t++ {
				//fmt.Println("Cheese 5")
				cardsToBeSaved = append(cardsToBeSaved, card)
			}
			y++
			x += numDuplicates
		}
	}

	if len(cardsToBeSaved) > 0 {
		writeGroupedOrSheet(cardsToBeSaved, writtenGroups)
	}
}

func writeGroupedOrSheet(cards []*image.RGBA, sheetNum int) {
	if grouped {
		writeSheet(cards, sheetNum)
	} else {
		writeIndividual(cards, sheetNum)
	}
}

//Writes out a whole sheet of images, that is 69, 10*7
func writeSheet(cards []*image.RGBA, sheetNum int) {
	exCard := cards[0]
	upLeft := image.Point{0, 0}
	iamgeRect := exCard.Bounds()
	cardWidth := iamgeRect.Dx()
	cardHeight := iamgeRect.Dy()
	totalwidth := iamgeRect.Dx() * 10
	totalhieght := iamgeRect.Dy() * 7

	lowRight := image.Point{totalwidth, totalhieght}

	//The blank image
	gridImage := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	sr := exCard.Bounds()
	for x, card := range cards {
		//Keep a loop of 10, then every 10, our divison adds one
		point := image.Point{(x % 10) * cardWidth, (x / 10) * cardHeight}
		r := image.Rectangle{point, point.Add(sr.Size())}
		draw.Draw(gridImage, r, card, sr.Min, draw.Src)
	}
	imgName := fmt.Sprintf("deck%d.png", sheetNum)
	f, _ := os.Create(imgName)
	png.Encode(f, gridImage)
}

func writeIndividual(cards []*image.RGBA, setNum int) {
	for x, card := range cards {
		cardName := fmt.Sprintf("card%d%d.png", setNum, x)
		f, _ := os.Create(cardName)
		png.Encode(f, card)
	}
}

//Write everything for one card
//data: string of fields
//locations: 2d array, with (x1,y1)(x2,y2) rectanlge text fitting
//background: background, we will deep copy here so dont create an extra copy for this
func writeCard(data []string, points [][4]int, background *image.RGBA) *image.RGBA {
	card := imageDeepCopy(background)
	for y, words := range data {
		addLabel(card, points[y], words)
		// Encode as PNG.
	}
	return card
}

func exportTogether() {

}

func exportIndividual() {

}

func countNumberOfcard(duplicates []int) int {
	total := 0
	for _, x := range duplicates {
		total += x
	}
	return total
}

func findDimensions(numberOfCards int) {
	MAX_WIDTH := 10
	MAX_HEIGHT := 7
	fmt.Println(MAX_HEIGHT + MAX_WIDTH)
}

func imageDeepCopy(img *image.RGBA) *image.RGBA {
	//Had bug before without star
	//Need to make cpy a actual aboject, not pointer to the img object
	cpy := *img

	cpy.Pix = make([]uint8, len(img.Pix))
	copy(cpy.Pix, img.Pix)
	return &cpy
}
