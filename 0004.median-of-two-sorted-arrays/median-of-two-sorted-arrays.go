package problem1

// FindMedianSortedArrays O(log(m+n))
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		indexResult1 int
		indexResult2 int
		lenAll       int
		countNow     int
		r1           int
		r2           int
		r1Find       bool
		r2Find       bool
		isEven       bool
	)
	lenAll = len(nums1) + len(nums2)
	if lenAll%2 == 0 {
		indexResult2 = lenAll / 2
		indexResult1 = indexResult2 - 1
		isEven = true
	} else {
		indexResult1 = (lenAll - 1) / 2
		indexResult2 = indexResult1
		isEven = false
	}

	var (
		iStart      int
		iEnd        int
		iMiddle     int
		jStart      int
		jEnd        int
		jMiddle     int
		jjs         int
		jje         int
		jDst        int
		outputCount int
	)
	iStart = 0
	iEnd = len(nums1) - 1
	jStart = 0
	jEnd = len(nums2) - 1
	for iStart <= iEnd {
		if (iStart+iEnd)%2 == 0 {
			iMiddle = (iStart + iEnd) / 2
		} else {
			iMiddle = (iStart + iEnd - 1) / 2
		}
		if iMiddle >= iEnd {
			outputCount++
			if outputCount > 1 {
				iMiddle = iEnd
				iStart = iEnd + 1
			}
		}
		jjs = jStart
		jje = jEnd
		jDst = 0
		for jjs <= jje {
			if (jjs+jje)%2 == 0 {
				jMiddle = (jjs + jje) / 2
			} else {
				jMiddle = (jjs + jje - 1) / 2
			}
			if nums2[jMiddle] > nums1[iMiddle] {
				if jje == jMiddle {
					break
				}
				jje = jMiddle
			} else if nums2[jMiddle] < nums1[iMiddle] {
				if jjs == jMiddle {
					jjs = jjs + 1
				} else {
					jjs = jMiddle
				}
			} else {
				jDst = jMiddle
				break
			}
		}
		if jDst == 0 {
			jDst = jjs
		}

		if (jDst + iMiddle) < indexResult1 {
			if iStart == iMiddle {
				iStart = iStart + 1
			} else {
				iStart = iMiddle
			}
			jStart = jjs
			countNow = iMiddle + 1
		} else if (jDst + iMiddle) > indexResult2 {
			iEnd = iMiddle
			jEnd = jje
		} else if (jDst + iMiddle) == indexResult1 {
			countNow = iMiddle + 1
			if !isEven {
				r2 = nums1[iMiddle]
				r1 = nums1[iMiddle]
				r1Find = true
				r2Find = true
				break
			} else {
				r1 = nums1[iMiddle]
				r1Find = true
				if r2Find {
					break
				} else {
					if iStart > iEnd {
						break
					}
					if iStart == iMiddle {
						iStart = iStart + 1
					} else {
						iStart = iMiddle
					}

				}
			}
		} else {
			countNow = iMiddle
			if !isEven {
				r2 = nums1[iMiddle]
				r1 = nums1[iMiddle]
				r1Find = true
				r2Find = true
				break
			} else {
				r2 = nums1[iMiddle]
				r2Find = true
				if r1Find {
					break
				} else {
					iEnd = iMiddle
				}
			}
		}
	}

	if !r1Find {
		if len(nums2) == 0 {
			r1 = nums1[iStart]
		} else {
			r1 = nums2[indexResult1-countNow]
		}

	}
	if !r2Find {
		if len(nums2) == 0 {
			r2 = nums1[iEnd]
		} else {
			r2 = nums2[indexResult2-countNow]
		}
	}
	return (float64(r1) + float64(r2)) / 2
}

// FindMedianSortedArraysX O(log(m+n))
func FindMedianSortedArraysX(nums1 []int, nums2 []int) float64 {
	var (
		indexResult1 int
		indexResult2 int
		lenAll       int
		countNow     int
		r1           int
		r2           int
		r1Find       bool
		r2Find       bool
		isEven       bool
	)
	lenAll = len(nums1) + len(nums2)
	if lenAll%2 == 0 {
		indexResult2 = lenAll / 2
		indexResult1 = indexResult2 - 1
		isEven = true
	} else {
		indexResult1 = (lenAll - 1) / 2
		indexResult2 = indexResult1
		isEven = false
	}

	var (
		iStart      int
		iEnd        int
		iMiddle     int
		jStart      int
		jEnd        int
		jMiddle     int
		jjs         int
		jje         int
		jDst        int
		outputCount int
	)
	iStart = 0
	iEnd = len(nums1) - 1
	jStart = 0
	jEnd = len(nums2) - 1
	for iStart < iEnd {
		if (iStart+iEnd)%2 == 0 {
			iMiddle = (iStart + iEnd) / 2
		} else {
			iMiddle = (iStart + iEnd - 1) / 2
		}
		if iMiddle+1 >= iEnd {
			outputCount++
			if outputCount > 1 {
				iMiddle = iEnd
				iStart = iEnd
			}
		}
		jjs = jStart
		jje = jEnd
		jDst = 0
		for jjs <= jje {
			if (jjs+jje)%2 == 0 {
				jMiddle = (jjs + jje) / 2
			} else {
				jMiddle = (jjs + jje - 1) / 2
			}
			if nums2[jMiddle] > nums1[iMiddle] {
				jje = jMiddle
				if jje == jMiddle {
					break
				}
			} else if nums2[jMiddle] < nums1[iMiddle] {
				if jjs == jMiddle {
					jjs = jjs + 1
				} else {
					jjs = jMiddle
				}
			} else {
				jDst = jMiddle
				break
			}
		}
		if jDst == 0 {
			jDst = jjs
		}

		if (jDst + iMiddle) < indexResult1 {
			iStart = iMiddle
			jStart = jjs
			countNow = iMiddle + 1
		} else if (jDst + iMiddle) > indexResult2 {
			iEnd = iMiddle
			jEnd = jje
		} else if (jDst + iMiddle) == indexResult1 {
			countNow = iMiddle + 1
			if !isEven {
				r2 = nums1[iMiddle]
				r1 = nums1[iMiddle]
				r1Find = true
				r2Find = true
				break
			} else {
				r1 = nums1[iMiddle]
				r1Find = true
				if r2Find {
					break
				} else {
					iStart = iMiddle
				}
			}
		} else {
			countNow = iMiddle
			if !isEven {
				r2 = nums1[iMiddle]
				r1 = nums1[iMiddle]
				r1Find = true
				r2Find = true
				break
			} else {
				r2 = nums1[iMiddle]
				r2Find = true
				if r1Find {
					break
				} else {
					iEnd = iMiddle
				}
			}
		}
	}

	if !r1Find {
		if len(nums2) == 0 {
			r1 = nums1[iStart]
		} else {
			r1 = nums2[indexResult1-countNow]
		}

	}
	if !r2Find {
		if len(nums2) == 0 {
			r2 = nums1[iEnd]
		} else {
			r2 = nums2[indexResult2-countNow]
		}
	}
	return (float64(r1) + float64(r2)) / 2
}

// FindMedianSortedArraysOm : two sort arrays media
func FindMedianSortedArraysOm(nums1 []int, nums2 []int) float64 {
	var (
		x          int
		y          int
		datalast   int
		datasecond int
		countData  int
		index      int
		isEven     bool
	)
	x = len(nums1)
	y = len(nums2)

	if (x+y)%2 == 0 {
		index = (x + y) / 2
		isEven = true
	} else {
		isEven = false
		index = (x + y - 1) / 2
	}
	var i, j = 0, 0
	for countData < index+1 {
		if i <= x-1 && j <= y-1 {
			if nums1[i] < nums2[j] {
				datasecond = datalast
				datalast = nums1[i]
				countData++
				i++
			} else {
				datasecond = datalast
				datalast = nums2[j]
				countData++
				j++
			}
		} else if i > x-1 && j <= y-1 {
			datasecond = datalast
			datalast = nums2[j]
			countData++
			j++
		} else if j > y-1 && i <= x-1 {
			datasecond = datalast
			datalast = nums1[i]
			countData++
			i++
		}
	}
	if isEven {
		return (float64(datasecond) + float64(datalast)) / 2
	}
	return float64(datalast)
}
