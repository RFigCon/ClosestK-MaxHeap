package main

import(
	"fmt"
)

type Point struct {

	x float64
	y float64

}

func (p Point) deltaToOrigin() float64{

	return p.x*p.x + p.y*p.y
	//No need to calculate the square root just to compare the distances
}

type Element struct {

	delta float64
	indx int64

}

type MaxHeap struct {

	size int64
	data[] *Element
	
}

func (heap *MaxHeap) swap( i, j int64){
	
	var temp *Element = heap.data[i]
	heap.data[i] = heap.data[j]
	heap.data[j] = temp

}

func (heap *MaxHeap) maxHeapify( indx int64 ){

	var largest, leftTreeIndx, rightTreeIndx int64
	var finished bool = false
	for !finished {

		largest = indx
		leftTreeIndx = indx*2 + 1
		rightTreeIndx = leftTreeIndx + 1

		if leftTreeIndx < heap.size && heap.data[indx].delta < heap.data[leftTreeIndx].delta {
			largest = leftTreeIndx
		}

		if rightTreeIndx < heap.size && heap.data[largest].delta < heap.data[rightTreeIndx].delta {
			largest = rightTreeIndx
		}

		if largest == indx {
			finished = true
		}else{
			heap.swap(indx, largest)
			indx = largest
		}

	}

}

func (heap *MaxHeap) buildHeap(){

	for i := heap.size/2 - 1; i>=0; i--{
		heap.maxHeapify( i )
	}

}

func (heap *MaxHeap) replaceMax( elm *Element ) {

	heap.data[0] = elm
	heap.buildHeap()

}

func (heap *MaxHeap) getMax() *Element{
	return heap.data[ 0 ]
}

var (
	heap MaxHeap = MaxHeap{size: 0}
	points[LEN] Point
)

const (
	LEN int64 = 10
	K int64 = 3
)

func main() {
	
	points = [LEN]Point{ {7, 2},
					{3, 20},
					{19, 14},
					{15, 17},
					{9, 18},
					{4, 19},
					{5, 6},
					{1, 8},
					{10, 11},
					{12, 13} }

	var temp[K] *Element
	var i int64
	for i = 0; i<K; i++ {
		temp[i] = &Element{ points[i].deltaToOrigin(), i }
	}

	heap = MaxHeap{ size: K, data: temp[:] }
	heap.buildHeap()

	var elm *Element
	for i = K; i<LEN; i++ {

		elm = &Element{points[i].deltaToOrigin(), i}
		if elm.delta < heap.getMax().delta {
			heap.replaceMax( elm )
		}

	}

	var indx int64
	for i = 0; i<K; i++ {
		indx = heap.data[i].indx
		fmt.Printf("%f %f\n", points[indx].x, points[indx].y)
	}
}