
# Redis Programming Training

## redis1
使用Redis存储和读取string类型数据


## redis2
使用Redis存储和读取Hash数据


## redis3

存储和解析json数据

## redis4

使用redis连接池，在init中初始化


# Scenario
Redis在不同场景的使用练习

## scenario1

Redis的List使用场景练习

在Web应用中，“列出最新的回复”之类的查询非常普遍，要输出这个顺序却不得不进行排序操作，这通常会带来性能和可扩展性问题。
类似的问题就可以用Redis来解决：
比如，我们的一个Web应用想要列出用户贴出的最新20条评论
假设数据库中的每条评论都有一个唯一的递增的ID字段，每次新评论发表时，我们会将它的ID添加到一个Redis列表


## scenario2

场景：排行榜应用，取TOP N操作
使用：Redis Sorted Set, 有序集合


最新N个数据是以某个条件为权重，比如按点赞的次数排序，这时候就需要sorted set
将你要排序的值设置成sorted set的score，将具体的数据设置成相应的value，每次只需要执行一条ZADD命令即可



