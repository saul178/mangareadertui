package components

/*
TODO:
opening a cbz file should not create a temp directory, it should read the cbz from memory and place it into a []byte buffer
store the image bytes into []byte slice, use bytes.reader
cache the image bytes into a temporary map[string]bytes so if the user goes between pages we are not wasting
cpu resources, if user moves on to another cbz we release the bytes from memory and clear the cache
CHECK the mimetype of the files before performing operations to read
*/
