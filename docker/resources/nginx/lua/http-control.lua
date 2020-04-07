local file = "/etc/lua/white_list"
local f = io.open(file, "rb")
local content = f:read("*all")
f:close()
ngx.header['Content-Type'] = 'text/plain';
for ip in string.gmatch(content, "[^\r\n]+") do
    if ngx.var.remote_addr == ip then
        ngx.exec('@meilisearch')
        return
    end
end

ngx.exit(ngx.HTTP_FORBIDDEN)