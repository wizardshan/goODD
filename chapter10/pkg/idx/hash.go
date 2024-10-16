package idx

import (
	"errors"
	"github.com/speps/go-hashids/v2"
)

const (
	HashIDCategoryUser int64 = iota
	HashIDCategoryArticle
)

var h *hashids.HashID

func init() {
	hd := hashids.NewData()
	hd.Salt = "elegantGo!@#"
	hd.MinLength = 10
	var err error
	h, err = hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
}

func Encode(id int64, c int64) string {
	hash, err := h.EncodeInt64([]int64{id, c})
	if err != nil {
		panic(err)
	}
	return hash
}

func Decode(hash string) (id int64, category int64, err error) {
	var numbers []int64
	if numbers, err = h.DecodeInt64WithError(hash); err != nil {
		return
	}

	if len(numbers) != 2 {
		err = errors.New("the numbers capacity expected value is 2")
		return
	}

	return numbers[0], numbers[1], nil
}
