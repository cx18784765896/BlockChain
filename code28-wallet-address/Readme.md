### 比特币钱包与地址

1. 通过钱包获取地址
2. 过程
    1. 公钥进行sha256再进行ripemd160得到公钥哈希
    2. 组成
        1. version：版本前缀，大小是一个字节，用来创造一个易于辨别的格式，“1”代表比特币地址
        2. pubKey hash：20字节，公钥hash
        3. checkSum：校验和，4字节，是添加到正在编码的数据的一端。校验和是通过pubKey哈希得到，用来检测输入时产生的错误
    3. （version+pubKeyHash+checkSum）得到两次哈希的地址
    4.  通过base58得到比特币地址