[![Logo](https://rclone.org/img/rclone-120x120.png)](https://rclone.org/)

[Website](https://rclone.org) |
[Documentation](https://rclone.org/docs/) |
[Contributing](CONTRIBUTING.md) |
[Changelog](https://rclone.org/changelog/) |
[Installation](https://rclone.org/install/) |
[Forum](https://forum.rclone.org/)
[G+](https://google.com/+RcloneOrg)

[![Build Status](https://travis-ci.org/artpar/rclone.svg?branch=master)](https://travis-ci.org/artpar/rclone)
[![Windows Build Status](https://ci.appveyor.com/api/projects/status/github/artpar/rclone?branch=master&passingText=windows%20-%20ok&svg=true)](https://ci.appveyor.com/project/artpar/rclone)
[![CircleCI](https://circleci.com/gh/artpar/rclone/tree/master.svg?style=svg)](https://circleci.com/gh/artpar/rclone/tree/master)
[![GoDoc](https://godoc.org/github.com/artpar/rclone?status.svg)](https://godoc.org/github.com/artpar/rclone) 

Rclone is a command line program to sync files and directories to and from

  * Amazon Drive
  * Amazon S3 / Dreamhost / Ceph / Minio / Wasabi
  * Backblaze B2
  * Box
  * Dropbox
  * FTP
  * Google Cloud Storage
  * Google Drive
  * HTTP
  * Hubic
  * Microsoft Azure Blob Storage
  * Microsoft OneDrive
  * Openstack Swift / Rackspace cloud files / Memset Memstore / OVH / Oracle Cloud Storage
  * pCloud
  * QingStor
  * SFTP
  * Webdav / Owncloud / Nextcloud
  * Yandex Disk
  * The local filesystem

Features

  * MD5/SHA1 hashes checked at all times for file integrity
  * Timestamps preserved on files
  * Partial syncs supported on a whole file basis
  * Copy mode to just copy new/changed files
  * Sync (one way) mode to make a directory identical
  * Check mode to check for file hash equality
  * Can sync to and from network, eg two different cloud accounts
  * Optional encryption (Crypt)
  * Optional FUSE mount

See the home page for installation, usage, documentation, changelog
and configuration walkthroughs.

  * https://rclone.org/

License
-------

This is free software under the terms of MIT the license (check the
COPYING file included in this package).
