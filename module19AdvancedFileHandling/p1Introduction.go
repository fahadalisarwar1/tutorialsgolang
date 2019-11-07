package main

// There are two main options for opening files, the os.Open() function, that simply accepts a filename
// and provides a ready-only file, and a more enhanced option, the os.OpenFile() function

// O_RDONLY Open the file read-only.
// O_WRONLY Open the file write-only.
// O_RDWR  Open the file read-write.
// O_APPEND Append data to the file when writing.
// O_CREATE Create a new file if none exists.
// O_EXCL Used with O_CREATE. The file must not exist.
// O_SYNC Open for Synchronous I/O.
// O_TRUNC If possible, truncate the file when opened.