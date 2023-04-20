package main

import (
	"fmt"
	//"git.dev.hnyunshangge.cn/backend/wx-push/dao/dao_openid"
)

var usersBroadcastBundle = 10000

func main() {
	// 群发图文到指定OpenID的用户
	len_users := 44300
	for start, end, round := 0, usersBroadcastBundle, 1; start < len_users; start, end, round = start+usersBroadcastBundle, end+usersBroadcastBundle, round+1 {
		if end > len_users {
			end = len_users
		}
		//msgID, msgDataID, err := callBroadcastApiRetryWrap(1_ctx, db, msg, users[start:end], 3)
		//if err == nil {
		//	dao_openid.SaveBroadcastOpenID(1_ctx, redis, msg.AppID, msgID, users[start:end])
		//}
		//logApiBroadcastResult(1_ctx, db, msg, task, round, end-start, msgID, msgDataID, err)
		fmt.Printf("start = %v end = %v\n", start, end)
	}
}