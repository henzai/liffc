# linectl

LINE LIFF, Things CLI tools in Go

## usage

```bash
# init
linectl init (LINE_ACCESS_TOKEN)

# add liff app
linectl liff add (URL)
# --description(-d) is remark of liff app
linectl liff add --description hoge (URL)
# --type(-t) full|tall|compact
linectl liff add --type tall (URL)
# --ble(-b) enable LINE Things
linectl liff add --ble (URL)

# delete liff app
linectl liff delete (LIFFID)
# delete all liff apps
linectl liff delete --all

# list all liff apps
linectl liff list

# update liff app
linectl liff update (LIFFID) (URL)

# send liff app URL to LINE
linectl liff send (LIFFID) (userID)
```
