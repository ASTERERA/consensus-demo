# consensus-demo
pow/pos共识算法demo

## pow算法
pow算法核心是通过大量的cpu运算，寻找一个满足某些条件的值（nonce）。
本demo中的算法比较简单，首选设置difficulty的值，然后求打包后的block的hash值，只要hash值满足条件：前`difficulty`个字符串的值为0，即为挖矿成功。

## pos算法
1. 本demo的pos算法，分为两个角色：validator是最后打包区块的账号，candidate是具有打包权限的候选者。
2. validator的选举是基于票数权重的，票数越多，被选中的概率就越大，挖的区块也就越多。
3. 可以通过http接口（`/api/pos/vote`）调用的方式，模拟为candidate投票，你会发现，获取票数越多的候选者，成为当前validator的几率越大

## 理解区块概念
1. 为便于理解区块及区块链的数据结构，本demo提供了一些相关的http接口，具体api在`server.Router()`方法内，每个接口实现都已加上了注释说明。
2. 启动程序后，可以看到控制台每隔一段时间会打印出新挖区块的高度和hash值。要想查看此区块的所有数据，可以调用`/api/blockchain/latest`
3. 调用`/api/blockchain`，可以打印当前已经挖出的所有区块

## 测试
1. 单元测试运行`consensus/pow/pow_test.go`和`consensus/pos/pos_test.go`里面的方法
2. 集成测试运行`tests/miner_test.go`里面的方法

## 尝试实现共识算法
在分支`question`内，你可以自己实现共识方法。全局查找`// TODO`，根据提示编写代码。
