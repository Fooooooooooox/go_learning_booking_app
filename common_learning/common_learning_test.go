package common_learning_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
)

func TestCommon(t *testing.T) {
	// 这个list是放了很多common.Address的实例 就像 Shapes := []Shape{c1, r1}
	list := []common.Address{
		common.HexToAddress("0x82a978b3f5962a5b0957d9ee9eef472ee55b42f1"),
		common.HexToAddress("0x7d577a597b2742b498cb5cf0c26cdcd726d39e6e"),
		common.HexToAddress("0xdceceaf3fc5c0a63d195d69b1a90011b7b19650d"),
		common.HexToAddress("0x598443f1880ef585b21f1d7585bd0577402861e5"),
		common.HexToAddress("0x13cbb8d99c6c4e0f2728c7d72606e78a29c4e224"),
		common.HexToAddress("0x77db2bebba79db42a978f896968f4afce746ea1f"),
		common.HexToAddress("0x24143873e0e0815fdcbcffdbe09c979cbf9ad013"),
		common.HexToAddress("0x10a1c1cb95c92ec31d3f22c66eef1d9f3f258c6b"),
		common.HexToAddress("0xe0fc04fa2d34a66b779fd5cee748268032a146c0"),
		common.HexToAddress("0x90f0b1ebbba1c1936aff7aaf20a7878ff9e04b6c"),
	}

	//这是一个新的地址 使用byte数组的形式来表示
	addressBytes := []byte{89, 132, 67, 241, 136, 14, 245, 133, 178, 31, 29, 117, 133, 189, 5, 119, 64, 40, 97, 229}

	var exists bool

	//遍历list中的所有地址
	// 将这些地址的byte和前面那个地址的byte一一比对 如果存在就把exists的值记为true
	for i, addr := range list {
		if bytes.Compare(addr.Bytes(), addressBytes) == 0 {
			fmt.Printf("found %s at index %v", addr.Hex(), i) // found 0x598443F1880Ef585B21f1d7585Bd0577402861E5 at index 3
			exists = true
			break
		}
	}

	//如果这个新地址在list中不存在 就把新地址添加到list
	if !exists {
		list = append(list, common.BytesToAddress(addressBytes))
	}

	spew.Dump(list)
}

//由上面这个例子可以知道 common.Address是geth库中的一种interface,你可以通过x.Hex()这种形式来调用这个interface支持的接口
