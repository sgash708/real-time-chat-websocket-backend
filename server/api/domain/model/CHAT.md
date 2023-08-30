# チャット処理の流れ

1. ReadLoopでwebsocketから新しいチャットを読み取る
2. 読み取りに成功すると、BroadCastChへ読み取った[]byteを送信する
3. BroadCastChへの送信をトリガーに、HubのbroadCastToAllClientが発火する。
4. broadCastToAllClientによって、HubのClientsフィールドが保持しているClient全てのsendChへ[]byteが送信される
5. sendChへの送信をトリガーに、WriteLoopのブロックが解除され、各ユーザーのブラウザにレスポンスが送られる 
