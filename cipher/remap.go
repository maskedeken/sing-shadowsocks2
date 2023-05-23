package cipher

import (
	"hash/crc32"

	mRand "math/rand"
)

func RemapToPrintable(input []byte) {
	const charSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\\\""
	seed := mRand.New(mRand.NewSource(int64(crc32.ChecksumIEEE(input))))
	for i := range input {
		input[i] = charSet[seed.Intn(len(charSet))]
	}
}
