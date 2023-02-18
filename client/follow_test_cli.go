package main

import (
	"context"
	"fmt"
	relation "github.com/ClubWeGo/relationmicro/kitex_gen/relation"
	relationService "github.com/ClubWeGo/relationmicro/kitex_gen/relation/relationservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Panic(err)
	}
	client, err := relationService.NewClient("relationservice", client.WithResolver(r))
	if err != nil {
		log.Panic(err)
	}

	followInfoResp, err := client.GetFollowInfoMethod(context.Background(), &relation.GetFollowInfoReq{MyUid: nil, TargetUid: 2009})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("followInfo无参", followInfoResp)
	myUid := new(int64)
	*myUid = 2006
	followInfoResp, err = client.GetFollowInfoMethod(context.Background(), &relation.GetFollowInfoReq{MyUid: myUid, TargetUid: 2009})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("followInfo有参", followInfoResp)

	//followList, err := client.GetFollowListReqMethod(context.Background(), &relation.GetFollowListReq{MyId: nil, TargetId: 2009})
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("无参followList")
	//for _, user := range followList.GetUserList() {
	//	fmt.Println(user)
	//}

	followList, err := client.GetFollowListReqMethod(context.Background(), &relation.GetFollowListReq{MyId: myUid, TargetId: 2009})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("有参followList")
	for _, user := range followList.GetUserList() {
		fmt.Println(user)
	}

	//followerList, err := client.GetFollowerListMethod(context.Background(), &relation.GetFollowerListReq{MyId: nil, TargetId: 2009})
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("无参followerList")
	//for _, user := range followerList.GetUserList() {
	//	fmt.Println(user)
	//}

	followerList, err := client.GetFollowerListMethod(context.Background(), &relation.GetFollowerListReq{MyId: myUid, TargetId: 2009})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("有参followerList")
	for _, user := range followerList.GetUserList() {
		fmt.Println(user)
	}


}