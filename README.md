# journ
journ is a journaling cli utility written in Go. Actual journals are done in `vim` editor (support for other editors coming later). All journals, after writing and quitting the editor, will be encrypted unless otherwise specified in a commandline flag. Encryption is currently symmetric, using [opengpg](https://godoc.org/golang.org/x/crypto/openpgp).

## main features ##
* encrypting/decrypting of journals
* automatic management of naming files and directory

## building journ ##
To build an executable: `go build -o journ`

## running journ ##
either of the following:
* `./journ`
* or `go run main.go`

## examples ##
`journ` will open vim editor. after writing and quitting, a file named `YYYY-MM-DDTHH:MM:SS-XX:XX.journal.pgp` will be in a directory named `journals` located in user's home directory.  
`journ -d /path/to/encrypted/journal.pgp` will decrypt the journal and write contents into another plaintext file  
`journ -s` will open editor for writing but will not encrypt its contents after
