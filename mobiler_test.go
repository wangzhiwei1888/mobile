package mobiler

import (
	"git.egret.io/go/assert"
	"testing"
)

func TestMobiler(t *testing.T) {
	mobiler := Mobiler{}

	err := mobiler.LoadData("./data.txt")

	if err != nil {
		assert.Panic(t, err, func() {})
	}

	mobile := mobiler.Mobile("18307400009")

	assert.Equal(t, "怀化", mobile.City)
	assert.Equal(t, "湖南", mobile.Province)
	assert.Equal(t, "移动", mobile.ISP)
	assert.Equal(t, "0745", mobile.AreaCode)
	assert.Equal(t, "418000", mobile.ZipCode)
	assert.Equal(t, "GSM/3G", mobile.Type)
	assert.Equal(t, "18307400009", mobile.Number)

	mobile = mobiler.Mobile("183****0009")

	assert.Equal(t, "移动", mobile.ISP)
	assert.Equal(t, "GSM/3G", mobile.Type)
	assert.Equal(t, "183****0009", mobile.Number)
	assert.Equal(t, "", mobile.City)
	assert.Equal(t, "", mobile.Province)
	assert.Equal(t, "", mobile.AreaCode)
	assert.Equal(t, "", mobile.ZipCode)

}
