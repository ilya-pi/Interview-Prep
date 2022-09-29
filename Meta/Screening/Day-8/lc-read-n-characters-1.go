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
	return func(buf []byte, n int) int {
		/*
		   To read n bytes by a window of 4, we will
		   read4 in a loop into 4 sized slice pointing
		   to underlying buf array, until we read less
		   then 4 bytes

		   So I need to have offset at which which to write to buf on each call and respect amount of bytes to read so far
		*/

		offset := 0
		totalRead := 0
		bf := make([]byte, 4)
		for totalRead < n {
			read := read4(bf)
			if read > n-totalRead {
				read = n - totalRead
			}
			copy(buf[offset:], bf[:read])
			totalRead += read
			offset += read
			if read < 4 {
				break
			}
		}
		return totalRead
	}
}
