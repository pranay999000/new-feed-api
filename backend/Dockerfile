FROM  golang:alpine
RUN mkdir -p /app/social-minor && cd /app/social-minor && \
	go mod init github.com/pranay999000/social-minor && \
	go get github.com/gin-contrib/cors  && \
	go get github.com/gin-contrib/sse  && \
	go get github.com/gin-gonic/gin  && \
	go get github.com/go-playground/locales  && \
	go get github.com/go-playground/universal-translator  && \
	go get github.com/go-playground/validator/v10  && \
	go get github.com/go-sql-driver/mysql  && \
	go get github.com/goccy/go-json  && \
	go get github.com/golang-jwt/jwt/v4  && \
	go get github.com/google/uuid  && \
	go get github.com/jinzhu/gorm  && \
	go get github.com/jinzhu/inflection  && \
	go get github.com/json-iterator/go && \
	go get github.com/leodido/go-urn  && \
	go get github.com/mattn/go-isatty  && \
	go get github.com/modern-go/concurrent && \
	go get github.com/modern-go/reflect2  && \
	go get github.com/pelletier/go-toml/v2  && \
	go get github.com/ugorji/go/codec  && \
	go get github.com/vence722/gin-middleware-bearer-token  && \
	go get golang.org/x/crypto  && \
	go get golang.org/x/net  && \
	go get golang.org/x/sys  && \
	go get golang.org/x/text  && \
	go get google.golang.org/protobuf  && \
	go get gopkg.in/yaml.v2

ADD . /app/social-minor
COPY . .
WORKDIR /app/social-minor
# RUN adduser -S -D -H -h /app appuser
# USER appuser
# RUN unset GOPATH
# RUN cd cmd/main
# RUN go install

# CMD ["/app/main"]
CMD ["go", "run", "/app/social-minor/cmd/main/main.go"]