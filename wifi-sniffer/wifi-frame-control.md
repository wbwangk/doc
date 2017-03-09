##wifi数据帧类型

控制帧 Frame Control 字段 2 个字节：

4bit(Subtype，小类)+2bit(Type，大类)+2bit(Protocol Version，默认为 00)，针对 Frame Control 的各 bit 位的说明如下：

### 管理帧：type为00
负责监督，用来加入或退出无线网络以及处理接入点之间关联的转移事宜。为了限制广播或组播管理帧所造成的副作用，收到管理帧后，必须加以查验。只有广播或者组播帧来自工作站当前所关联的 BSSID 时，它们才会被送至 MAC 管理层。唯一例外是beacon 帧  
此时各 Subtype 的值如下：  

    值        |      含义
--------------|-----------------------
    0000      |     Association request（连接要求）
    0001      |     Association response（连接应答）
    0010      |     Reassociation request（重新连接要求）
    0011      |     Reassociation response（重新连接应答）
    0100      |     Probe request（探查要求）
    0101      |     Probe response（探查应答）
    1000      |     Beacon（导引信号）
    1001      |     Announcement  traffic  indication  message (ATIM)    
    1010      |     Disassociation（解除连接）
    1011      |     Authentication（身份验证）
    1100      |     Deauthentication（解除认证）

### 控制帧：type为01
   
   subType       |          含义
-----------------|--------------------------
   1010          |         Power Save‐Poll（省电模式－轮询）
   1011          |         RTS（请求发送）
   1100          |         CTS（允许发送）
   1101          |         ACK（应答）
   1110          |         CF‐End（免竞争期间结束）
   1111          |         CF‐End（免竞争期间结束）+CF‐Ack（免竞争期间回应）
   1001          |         块回应


### 数据帧：type为10
subType       |          含义
--------------|---------------------------
0000          |          Data（数据）(0x08)
0001          |          Data+CF‐Ack
0010          |          Data+CF‐Poll (0x28)
0011          |          Data+CF‐Ack+CF‐Poll
0100          |          Null data (无数据：未发送数据)(0x48)
0101          |          CF‐Ack (未发送数据)
0110          |          CF‐Poll (未发送数据)
0111          |          Data＋CF‐Ack+CF‐Poll
1000          |          QoS Data【注 c】 (0x88)
1001          |          QoS Data + CF‐Ack
1010          |          QoS Data + CF‐Poll
1011          |          QoS Data + CF‐Ack + CF‐Pol
1100          |          QoS Null (未发送数据)
1101          |          QoS CF‐Ack (未发送数据)
1110          |          QoS CF‐Poll (未发送数据)
1111          |          QoS CF‐Ack+CF‐Poll （未发送数据)
