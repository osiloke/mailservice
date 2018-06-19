# Mailsender

This is a service for sending mails through various smtp services. It uses machinery for handling new requests to send mails.

## Getting Started

<!-- Make sure you have re -->

### Prerequisities

In order to run this container you'll need docker installed.

- [Windows](https://docs.docker.com/windows/started)
- [OS X](https://docs.docker.com/mac/started/)
- [Linux](https://docs.docker.com/linux/started/)

### Usage

#### Container Parameters

Run the service

```shell
docker run mailsender
```

<!--
One example per permutation

```shell
docker run give.example.org/of/your/container:v0.2.1
``` -->

You can start a shell in the service by running the command below

```shell
docker run mailsender:latest bash
```

#### Environment Variables

- `REDIS_URI` - Redis uri - default {localhost:6379}
- `QUEUE` - Name of the default redis queue - default {machinery_tasks}
- `NAME` - name of this service

<!-- #### Volumes

* `/your/file/location` - File location -->

<!-- #### Useful File Locations

* `/some/special/script.sh` - List special scripts

* `/magic/dir` - And also directories -->

## Built With

<!-- * List the software v0.1.3
* And the version numbers v2.0.0
* That are in this container v0.3.2 -->

## Find Us

- [GitHub](https://github.com/osiloke/mailsender)
  <!-- * [Quay.io](https://quay.io/repository/your/docker-repository) -->

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the
[tags on this repository](https://github.com/osiloke/mailsender/tags).

## Authors

- **Osiloke Emoekpere** - _Initial work_ - [Osiloke](http://osiloke.com)

See also the list of [contributors](https://github.com/your/repository/contributors) who
participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- docker
