package main

import (
	"fmt"
	"sort"
)

/*
requirement clarification
	1. box: contains boxes and items
	2. item: a product with attributes "name"
	3. retrieve all items within a box
	4. retrieve all items and boxes within a box
	5. put and remove boxes and items into a box
	6. no need to handle concurrent requests
	7. assume all inputs and operations are valid, no need to throw error
	8. output can be of different sequence
*/

type Item interface {
	GetID() int64
	GetName() string
}

type Box interface {
	GetID() int64
	GetBoxes() []Box
	GetItems() []Item
	AddItem(Item)
	RemoveItem(int64)
	AddBox(Box)
	RemoveBox(int64)
}

// Item implementation
type item struct {
	id int64
	name string
}

func (i *item) GetID() int64 {
	return i.id
}

func (i *item) GetName() string {
	return i.name
}

func NewItem(id int64, name string) Item {
	return &item{
		id: id,
		name: name,
	}
}

// Box implementation
type box struct {
	id int64
	boxes map[int64]Box
	items map[int64]Item
}

func (b *box) GetID() int64 {
	return b.id
}

func (b *box) GetBoxes() []Box {
	boxes := []Box{}
	for _, subBox := range b.boxes {
		boxes = append(boxes, subBox)
	}
	return boxes
}

func (b *box) GetItems() []Item {
	items := []Item{}
	for _, subItem := range b.items {
		items = append(items, subItem)
	}
	return items
}

func (b *box) AddItem(newItem Item) {
	b.items[newItem.GetID()] = newItem
}

func (b *box) RemoveItem(id int64) {
	delete(b.items, id)
}

func (b *box) AddBox(newBox Box) {
	b.boxes[newBox.GetID()] = newBox
}

func (b *box) RemoveBox(id int64) {
	delete(b.boxes, id)
}

func NewBox(id int64, boxes []Box, items []Item) Box {
	newBox := &box{
		id: id,
		boxes: make(map[int64]Box),
		items: make(map[int64]Item),
	}
	for _, b := range boxes {
		newBox.AddBox(b)
	}
	for _, i := range items {
		newBox.AddItem(i)
	}
	return newBox
}

// GetAllItemsInBox returns all the items in the input box and its sub boxes
func GetAllItemsInBox(targetBox Box) []Item {
	allItems := []Item{}
	var dfs func(currentBox Box)
	dfs = func(currentBox Box) {
		if currentBox == nil {
			return
		}
		allItems = append(allItems, currentBox.GetItems()...)
		for _, nextBox := range currentBox.GetBoxes() {
			dfs(nextBox)
		}
	}
	dfs(targetBox)
	return allItems
}

func main() {
	floss := NewItem(0, "floss")
	toothbrush := NewItem(1, "toothbrush")
	toothpaste := NewItem(2, "toothpaste")

	boxC := NewBox(0, nil, []Item{floss})
	boxB := NewBox(1, nil, []Item{toothbrush, toothpaste})
	boxA := NewBox(2, []Box{boxB, boxC}, nil)

	allItems := GetAllItemsInBox(boxA)
	for _, itemInBox := range allItems {
		fmt.Println(itemInBox.GetName())
	}
}


type operation struct {
    time int
    enter bool
}

func minMeetingRooms(intervals [][]int) int {
    n := len(intervals)
    operations := make([]operation, 2*n)
    for i, interval := range intervals {
        operations[2*i] = operation{interval[0], true}
        operations[2*i+1] = operation{interval[1], false}
    }
    sort.Slice(operations, func(i, j int) bool {
        if operations[i].time == operations[j].time {
            return !operations[i].enter
        }
        return operations[i].time < operations[j].time
    })

    roomCount := 0
    maxRoomCount := 0
    for _, op := range operations {
        if op.enter {
            roomCount++
            if roomCount > maxRoomCount {
                maxRoomCount = roomCount
            }
        } else {
            roomCount--
        }
    }

    return maxRoomCount
}