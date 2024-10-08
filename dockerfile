FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /app/main

# 添加应用可执行文件，并设置执行权限
ADD ./main   $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加静态资源文件
# ADD resource $WORKDIR/resource

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ["./main"]