FROM scratch

WORKDIR $GOPATH/src/rehabilitation_prescription
COPY . $GOPATH/src/rehabilitation_prescription

EXPOSE 8000
CMD ["./rehabilitation"]


