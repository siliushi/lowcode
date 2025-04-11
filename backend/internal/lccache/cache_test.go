package lccache

import (
	"fmt"
	"testing"
	"time"

	cache "github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

const (
	KSchool string = "school"
	KAge    string = "age"
)

func GetSchool(cache *cache.Cache) string {
	schObj, found := cache.Get(KSchool)
	var schoolName string
	if found {
		schoolName = schObj.(string)
	}
	fmt.Println("schoolName = ", schoolName)
	return schoolName
}

func GetAge(cache *cache.Cache) int {
	ageObj, found := cache.Get(KAge)
	var age int
	if found {
		age = ageObj.(int)
	}
	fmt.Println("age = ", age)
	return age
}

func TestGoCache(t *testing.T) {
	t.Skip()
	assert := assert.New(t)
	// 默认过期时间: 5*time.Minute
	// 清理过期的缓存对象: 10*time.Minute
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set(KSchool, "hangzhou-001", time.Second*3)
	c.Set("age", 42, cache.NoExpiration)
	assert.Equal("hangzhou-001", GetSchool(c))
	assert.Equal(42, GetAge(c))
	time.Sleep(time.Second * 5)
	assert.Equal("", GetSchool(c))
	assert.Equal(42, GetAge(c))
}
