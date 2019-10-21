function readAll(file)
    local f = assert(io.open(file, "rb"))
    local content = f:read("*all")
    f:close()
    return content
end

wrk.method = "POST"
wrk.body = readAll("/var/local/happyconv/books.xml")
wrk.headers["Content-Type"] = "application/xml"
