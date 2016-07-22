# gzipresponse
gzip HTTP response writer for Go.

This package contains just one function called `Write`, which you use to write your data into an `http.ResponseWriter`.

We created this as an alternative to http://github.com/NYTimes/gziphandler, because it seems impossible to use that handler in combination with http://github.com/julienschmidt/httprouter.
