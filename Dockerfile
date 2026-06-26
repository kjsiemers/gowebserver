# This is a comment

# Use a lightweight debian os
# as the base image
FROM debian:stable-slim
COPY goserver /bin/goserver

# execute the web server
# command when the container runs
ENV PORT=8991
CMD ["/bin/goserver"]
