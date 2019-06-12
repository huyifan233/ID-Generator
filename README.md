# ID-Generator

一款基于区块链(以太坊)的ID生成器

An ID generator based on blockchain(Ethereum)

## 用法
- 按照conf/conf.yaml中的注释提示，修改conf/conf.yaml中的配置。
- 在urls.go中启动程序。
- 访问http://xxx.xxx.xxx.xxx:port/ids 获得Id的最大值。比如访问获得值为1000(取决于step的值)，那么可用的Id范围为(1000-step+1) —— 1000，得出Id范围就是1 —— 1000。如果获得2000，那么同理你的可用Id范围为1001 —— 2000。Id生成的规律为step,2*step,3*step ..... n*step。
