# fileserver

## POST

### /file/create

curl -s -X POST http://localhost:3000/file/create -d 'uuid=uuid0001&name=额测试解耦&desc=这是一个ss个&url=http://127.dkjf/?dfdk&hash=sdfakdfajsdfk'

### /file/:uuid

http://localhost:3000/file/uuid0001

## upload file

http://localhost:3000/upload

    <form action="http://123.206.29.15:3000/upload" method="post" enctype="multipart/form-data">
        UUID:
        <input type="text" name="uuid">
        <br> 文件名称:
        <input type="text" name="name">
        <br> 描述:
        <input type="text" name="desc">
        <br> Files:
        <input type="file" name="file">
        <br>
        <br>
        <input type="submit" value="Submit">
    </form>

