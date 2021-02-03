local arg = ngx.req.get_uri_args()
for k, v in pairs(arg) do
     ngx.say(k .. "-" ..v)
end
