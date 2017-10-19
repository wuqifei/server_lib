package libzookeeper

import (
	"fmt"
	"path"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type Option struct {
	// zk的连接超时时间
	Timeout time.Duration
	// zkserver的地址，例如:localhost:2381,localhost:2382/test
	// 包括了topic和地址
	Server string

	//zookper的输出日志
	Logger zk.Logger

	addrs  []string
	chroot string
}

// zookeeper实例
type ZooKeeper struct {
	conn   *zk.Conn
	option *Option
}

// 生成一个zk
func NewZK(option *Option) (*ZooKeeper, error) {
	if option == nil {
		panic(fmt.Errorf("zk:config is nil"))
	}

	var err error
	addrs, chroot := ParseZKAddrString(option.Server)
	zookeeper := &ZooKeeper{}
	option.addrs = addrs
	option.chroot = chroot

	zookeeper.option = option

	if option.Logger != nil {

		logOption := func(c *zk.Conn) {
			c.SetLogger(option.Logger)
		}
		if zookeeper.conn, _, err = zk.Connect(addrs, option.Timeout, logOption); err == nil {
			return zookeeper, nil
		}
	} else {
		if zookeeper.conn, _, err = zk.Connect(addrs, option.Timeout); err == nil {
			return zookeeper, nil
		}
	}

	return nil, err
}

// 继承closer接口
func (z *ZooKeeper) Close() error {
	z.conn.Close()
	return nil
}

// 返回zookeeper的所有节点信息
func (z *ZooKeeper) Nodes(node string) (map[string]string, error) {
	if node[0] != '/' {
		node = "/" + node
	}
	root := fmt.Sprintf("%s%s", z.option.chroot, node)
	children, _, err := z.conn.Children(root)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for _, child := range children {
		if err != nil {
			return nil, err
		}
		val, _, err := z.conn.Get(path.Join(root, child))
		if err != nil {
			return nil, err
		}

		result[child] = string(val)
	}
	return result, nil
}

// 节点是否存在
func (z *ZooKeeper) Exists(node string) (ok bool, err error) {
	if node[0] != '/' {
		node = "/" + node
	}
	root := fmt.Sprintf("%s%s", z.option.chroot, node)
	ok, _, err = z.conn.Exists(root)
	return
}

// 循环删除
func (z *ZooKeeper) DeleteRecursive(node string) (err error) {
	if node[0] != '/' {
		node = "/" + node
	}

	root := fmt.Sprintf("%s%s", z.option.chroot, node)
	return z.deleteRecursive(root)
}

func (z *ZooKeeper) deleteRecursive(node string) (err error) {
	children, stat, err := z.conn.Children(node)
	if err == zk.ErrNoNode {
		return nil
	} else if err != nil {
		return
	}

	for _, child := range children {
		if err = z.deleteRecursive(path.Join(node, child)); err != nil {
			return
		}
	}
	return z.conn.Delete(node, stat.Version)
}

// 循环建立node
func (z *ZooKeeper) MkdirRecursive(node string) (err error) {
	if node[0] != '/' {
		node = "/" + node
	}

	root := fmt.Sprintf("%s%s", z.option.chroot, node)
	return z.mkdirRecursive(root)
}
func (z *ZooKeeper) mkdirRecursive(node string) (err error) {
	parent := path.Dir(node)
	if parent != "/" {
		if err = z.mkdirRecursive(parent); err != nil {
			return
		}
	}
	_, err = z.conn.Create(node, nil, 0, zk.WorldACL(zk.PermAll))
	if err == zk.ErrNodeExists {
		err = nil
	}
	return
}

// 在节点创建一个新的value，如果已经创建了则返回错误，如果是临时节点则填入true
func (z *ZooKeeper) Create(node string, value []byte, ephemeral bool) (err error) {
	if node[0] != '/' {
		node = "/" + node
	}

	root := fmt.Sprintf("%s%s", z.option.chroot, node)
	if err = z.mkdirRecursive(path.Dir(root)); err != nil {
		return
	}
	flag := int32(0)
	if ephemeral {
		flag = zk.FlagEphemeral
	}
	_, err = z.conn.Create(root, value, flag, zk.WorldACL(zk.PermAll))
	return
}
