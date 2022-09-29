/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	overflow := make([]byte, 4)
	overflowCount := 0
	return func(buf []byte, n int) int {
		buf4 := make([]byte, 4)
		totalRead := 0
		offset := 0
		if overflowCount > n {
			copy(buf, overflow[:n])
			copy(overflow, overflow[n:])
			overflowCount -= n
			totalRead = n
		} else {
			copy(buf, overflow[:overflowCount])
			offset = overflowCount
			totalRead = overflowCount
			overflowCount = 0
		}
		for totalRead < n {
			read := read4(buf4)
			if read > n-totalRead {
				// record overflow
				overflowCount = read - (n - totalRead)
				read = n - totalRead
				copy(overflow, buf4[read:])
			}
			copy(buf[offset:], buf4[:read])
			totalRead += read
			offset += read
			if read < 4 {
				break
			}
		}
		return totalRead
	}
}
