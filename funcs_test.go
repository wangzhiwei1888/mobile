package mobile

import (
	"git.egret.io/go/assert"
	"testing"
)

var (
	mobile       = "18307400009"
	mobileEase   = "183****0009"
	mobileString = "18307400009 还有 13973333488 还有 183****8188"
)

func TestIsMobile(t *testing.T) {
	assert.Equal(t, true, IsMobile(mobile, true))
	assert.Equal(t, true, IsMobile(mobile, false))
	assert.Equal(t, true, IsMobile(mobileEase, false))
	assert.Equal(t, false, IsMobile(mobileEase, true))
}

func TestHasMobile(t *testing.T) {
	assert.Equal(t, true, HasMobile(mobile, true))
	assert.Equal(t, true, HasMobile(mobile, false))
	assert.Equal(t, true, HasMobile(mobileEase, false))
	assert.Equal(t, false, HasMobile(mobileEase, true))
	assert.Equal(t, true, HasMobile(mobileString, true))
	assert.Equal(t, true, HasMobile(mobileString, false))
}

func TestCountMobile(t *testing.T) {
	assert.Equal(t, 1, CountMobile(mobile, true))
	assert.Equal(t, 1, CountMobile(mobile, false))
	assert.Equal(t, 1, CountMobile(mobileEase, false))
	assert.Equal(t, 0, CountMobile(mobileEase, true))
}

func TestGetMobiles(t *testing.T) {
	assert.Equal(t, 2, len(GetMobiles(mobileString, true)))
	assert.Equal(t, 3, len(GetMobiles(mobileString, false)))
	assert.Equal(t, 1, len(GetMobiles(mobile, true)))
	assert.Equal(t, 1, len(GetMobiles(mobile, false)))
	assert.Equal(t, 0, len(GetMobiles(mobileEase, true)))
	assert.Equal(t, 1, len(GetMobiles(mobileEase, false)))
	assert.Equal(t, "18307400009", GetMobiles(mobile, true)[0])
	assert.Equal(t, "18307400009", GetMobiles(mobile, false)[0])
	assert.Equal(t, "18307400009", GetMobiles(mobileString, true)[0])
	assert.Equal(t, "18307400009", GetMobiles(mobileString, false)[0])
	assert.Equal(t, "13973333488", GetMobiles(mobileString, true)[1])
	assert.Equal(t, "13973333488", GetMobiles(mobileString, false)[1])
	assert.Equal(t, "183****8188", GetMobiles(mobileString, false)[2])
}
