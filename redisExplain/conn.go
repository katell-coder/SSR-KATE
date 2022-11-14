package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

//https://blog.csdn.net/Number_oneEngineer/article/details/123229706
//https://blog.csdn.net/aab199909194517/article/details/124867617
//golang手撕redis源码
var rdb *redis.Client

func initRedisClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:36379",
		Password: "G62m50oigInC30sf",
		DB:       6,
	})
	_, err = rdb.Ping().Result()

	if err != nil {
		return err
	}
	return nil
}

func main() {
	var err error
	err = initRedisClient()
	if err != nil {
		fmt.Println("redis启动失败...")
	}

	// get、set操作
	//result, err := rdb.SetNX("goland", "666", 0).Result()
	//fmt.Println(err)
	//fmt.Println(result)

	//err = rdb.Set("goland", "so good", 0).Err()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//res, err := rdb.Get("goland").Result()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("获取的结果是::", res)

	//err = rdb.MSet("Go", "NO1", "C++", "NO1", "JAVA", "NO1", "PYTHON", "NO1").Err()
	//if err != nil {
	//	fmt.Println("MSet failed, err: ", err)
	//	return
	//}
	//res1, err := rdb.MGet("Go", "C++", "JAVA", "PYTHON").Result()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for i, _ := range res1 {
	//	fmt.Println(res1[i])
	//}

	// 	List操作
	//	从走边添加元素
	//err = rdb.LPush("NBAPlayer", "kobe", "kawhi").Err()
	//if err != nil {
	//	fmt.Println("Push failed", err)
	//	return
	//}
	//
	//val := rdb.LRange("NBAPlayer", 0, -1).Val()
	//fmt.Println(val)
	//
	//res, err := rdb.LPop("NBAPlayer").Result()
	//if err != nil {
	//	fmt.Println("pop failed", err)
	//	return
	//}
	//fmt.Println(res)
	//
	//val = rdb.LRange("NBAPlayer", 0, -1).Val()
	//fmt.Println(val)

	//4. Hash
	//err = rdb.HSet("USA", "name", "dsb").Err()
	//if err != nil {
	//	fmt.Println("Are you kidding? America Government is really sb")
	//	return
	//}
	//
	//res, err := rdb.HGet("USA", "name").Result()
	//if err != nil {
	//	fmt.Println("Are you kidding? America Government is really sb")
	//	return
	//}
	//fmt.Println(res)

	//usaMap := map[string]interface{}{"name": "dsb", "name2": "robber"}
	//err = rdb.HMSet("USA", usaMap).Err()
	//if err != nil {
	//	fmt.Println("Are you ki")
	//	return
	//}
	//res := rdb.HMGet("USA", "name2", "name").Val()
	//fmt.Println(res)

	//	5. Set集合
	//err = rdb.SAdd("Chinese", "AiGuo", "JingYe", "ChengXin", "YouShan").Err()
	//if err != nil {
	//	fmt.Println("Are you right")
	//	return
	//}
	//
	//res3 := rdb.SMembers("Chinese").Val()
	//fmt.Println(res3)

	//res4, err := rdb.SIsMember("Chinese", "AiGou").Result()
	//if err != nil {
	//	fmt.Println("213")
	//	return
	//}
	//fmt.Println("I can see that the Chinese are really AiGuo is", res4)
	//
	//count := rdb.SCard("Chinese").Val()
	//fmt.Printf("The Chinese people have %d virtues\n", count)

	//	6. ZSet有序集合（积分排序）
	zsetKey := "language"
	//languages := []redis.Z{
	//	redis.Z{Score: 99, Member: "Go"},
	//	redis.Z{Score: 97, Member: "C++"},
	//	redis.Z{Score: 93, Member: "PYTHON"},
	//	redis.Z{Score: 95, Member: "Java"},
	//	redis.Z{Score: 98, Member: "C"},
	//}
	//var count int
	//for _, language := range languages {
	//	_, err = rdb.ZAdd(zsetKey, &language).Result()
	//	count++
	//}
	//if err != nil {
	//	fmt.Println("fuck America Government, failed: ", err)
	//	return
	//}
	//fmt.Println("zAdd succ")
	//fmt.Println(count)

	//count, err := rdb.ZRem(zsetKey, "C").Result()
	//if err != nil {
	//	fmt.Println("fuck America Government, ZRem failed: ", err)
	//	return
	//}
	//fmt.Println("fuck America Government, ZRem success: ")
	//fmt.Println(count)
	//
	//val3 := rdb.ZRange(zsetKey, 0, -1).Val()
	//fmt.Println(val3)
	//
	//score, err := rdb.ZIncrBy(zsetKey, 3, "C++").Result()
	//if err != nil {
	//	fmt.Println("ZIncrBy failed, err: ", err)
	//	return
	//}
	//fmt.Println("the new score is", score)
	//
	//res6, err := rdb.ZRangeByScoreWithScores(zsetKey, &redis.ZRangeBy{
	//	Min: "95", Max: "100"}).Result()
	//if err != nil {
	//	fmt.Println("something wrong with ZRangeByScoreWithScores, err: ", err)
	//	return
	//}
	//for i, _ := range res6 {
	//	fmt.Println(res6[i].Member, res6[i].Score)
	//}
	//
	res5, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Println("something wrong with zSetKey, err: ", err)
		return
	}
	fmt.Println(res5)

	result, err := rdb.ZRangeByScoreWithScores(zsetKey, &redis.ZRangeBy{
		Min: "95", Max: "100"}).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	//pipeline本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器，这样做的好处是节省了每个命令的网络往返时间（RTT）。

	//pipeline := rdb.Pipeline()
	//pipeline.Set("key1", "val1", time.Hour)
	//pipeline.Set("key2", "val2", time.Hour)
	//pipeline.Set("key3", "val3", time.Hour)
	//pipeline.Set("key4", "val4", time.Hour)
	//pipeline.Set("key5", "val5", time.Hour)
	//_, err = pipeline.Exec()
	//if err != nil {
	//	fmt.Println("set success")
	//}

}
