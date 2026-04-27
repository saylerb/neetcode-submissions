// for this problem
// when you get, you get the most recent value
// and the most recent timestamp for that key is less than or equal to the timestamp

// Naive approach
// store a list of Data for each key
// look at the last item (most recent), if it is <= than the queried timestamp, return it

// if it's not less than  for example we have [2,3,4] and 3 is queried, we need to do a lookup

// the most efficient way to do this lookup would be to do a binary search
// over the array, assuming that array is always sorted by timestamp

// since calls to set are in increasing order specifically, we can assume that 
// it will be in order, as along as we always append

type Data struct {
	value string
	timestamp int
}

type TimeMap struct {
	data map[string][]Data
}

func Constructor() TimeMap {
	data := make(map[string][]Data)
	return TimeMap{ data }
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	data := Data{ value, timestamp }
	this.data[key] = append(this.data[key], data)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	items := this.data[key]

	return binarySearch(items, timestamp)
}

func binarySearch(items []Data, timestamp int) string {
	n := len(items)
	if n == 0 {
		return ""
	}

	l, r := 0, n - 1
	// edge case == l is the item
	// edge case, l > item (the whole range is after the timestamp)
	// edge case, r < item (the whole range is before the timestamp)
	// if the whole reange is before the timestamp, we just take the last item

	if items[l].timestamp >= timestamp { // bug: if the whole range is before the last item
		if items[l].timestamp == timestamp { // bug, didn't compare timestamps
			return items[l].value
		} 
		return ""
	}
	// do we need to return the last item, in the case that r < timestamp?

	for r - l > 1 {
		m := (r + l) / 2

		if items[m].timestamp >= timestamp {
			r = m
		} else {
			l = m
		}
	}
	if items[r].timestamp <= timestamp {
		return items[r].value
	} else {
		return items[l].value
	}
}
