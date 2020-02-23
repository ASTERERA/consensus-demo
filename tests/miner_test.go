package tests

import (
	"consensus-demo/miner"
	"consensus-demo/node"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestMine(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testMine("pow", t)
	testMine("pos", t)
}

func testMine(consensus string, t *testing.T) {
	t.Logf("begin to test miner, current consensus is: %s\n", consensus)
	n, err := node.New(node.Config{
		Name: "test-node",
		Difficulty: 1,
		Consensus: consensus,
	})
	if err != nil {
		t.Fatalf("create node error: %s\n", err.Error())
	}
	go func() {
		n.Start()
	}()

	time.Sleep((miner.MineInterval*1.5) * time.Second) // waiting for miner to generate a block
	height, err := httpGet("http://127.0.0.1:8080/api/blockchain/height")
	if err != nil {
		t.Fatalf("http request error: %s\n", err.Error())
	}
	expected := "1"
	if height != expected {
		t.Fatalf("wrong blockchain height, expect: %s, got: %s\n", expected, height)
	}
	n.Stop()

	t.Log("test miner success")
}

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	sbody := string(body)
	// fmt.Println(sbody)
	return strings.TrimRight(sbody, "\n"), nil
}
