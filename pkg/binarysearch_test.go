package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)
	{
		arr := []int{1, 2, 4, 5}
		assert.Equal(-1, BinarySearch(arr, 0))
		assert.Equal(0, BinarySearch(arr, 1))
		assert.Equal(1, BinarySearch(arr, 2))
		assert.Equal(-1, BinarySearch(arr, 3))
		assert.Equal(2, BinarySearch(arr, 4))
		assert.Equal(3, BinarySearch(arr, 5))
		assert.Equal(-1, BinarySearch(arr, 6))
	}
	{
		arr := []int{1, 2, 4, 5, 6}
		assert.Equal(-1, BinarySearch(arr, 0))
		assert.Equal(0, BinarySearch(arr, 1))
		assert.Equal(1, BinarySearch(arr, 2))
		assert.Equal(-1, BinarySearch(arr, 3))
		assert.Equal(2, BinarySearch(arr, 4))
		assert.Equal(3, BinarySearch(arr, 5))
		assert.Equal(4, BinarySearch(arr, 6))
		assert.Equal(-1, BinarySearch(arr, 7))
	}
}

func TestBinarySearch1(t *testing.T) {
	assert := assert.New(t)
	{
		arr := []int{1, 1, 1, 2, 2, 2, 4, 4, 4, 5, 5, 5}
		assert.Equal(0, BinarySearch1(arr, 0))
		assert.Equal(3, BinarySearch1(arr, 1))
		assert.Equal(6, BinarySearch1(arr, 2))
		assert.Equal(6, BinarySearch1(arr, 3))
		assert.Equal(9, BinarySearch1(arr, 4))
		assert.Equal(12, BinarySearch1(arr, 5))
		assert.Equal(12, BinarySearch1(arr, 6))
	}
}

func TestBinarySearch2(t *testing.T) {
	assert := assert.New(t)
	{
		arr := []int{1, 1, 1, 2, 2, 2, 4, 4, 4, 5, 5, 5}
		assert.Equal(-1, BinarySearch2(arr, 0))
		assert.Equal(-1, BinarySearch2(arr, 1))
		assert.Equal(2, BinarySearch2(arr, 2))
		assert.Equal(5, BinarySearch2(arr, 3))
		assert.Equal(5, BinarySearch2(arr, 4))
		assert.Equal(8, BinarySearch2(arr, 5))
		assert.Equal(11, BinarySearch2(arr, 6))
	}
}
