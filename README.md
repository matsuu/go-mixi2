# go-mixi2

Tool for posting to mixi2.

## Prepare

```
go install github.com/matsuu/go-mixi2/cmd/mixi2_post@latest
```

## Usage

* Access https://mixi.social/ in your browser and log in.
* After logging in, open the Developer Tools (usually by pressing F12) and navigate to the Network tab.
* While keeping the Network tab open, make a test post.
* Find and select the CreatePost request in the Network tab list.
* In the Request Headers section for the selected CreatePost request, find the following:
    * The value of auth\_token within the Cookie header.
    * The value of the X-Auth-Key header.
    * The value of the X-Mercury-User-Agent header.
* Pass the values you have identified as follows:

```
mixi2_post -token="your_auth_token" -key="your_auth_key" -ua="your_user_agent" -text="your_soul"
```
