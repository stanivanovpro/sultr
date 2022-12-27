FROM mongo:6.0
WORKDIR /app
COPY fixtures /home/fixtures
COPY bin /home/bin
COPY . /app
EXPOSE 27017
CMD ["mongod", "--bind_ip_all"]
RUN /home/bin/install-fixtures.sh