## ghsts
cross-cutting github repo setting tool

## How to use

```
$ go get -u github.com/wasanx25/ghsts
$ cd /path/to/ghsts
$ GO111MODULE=on go install
$ export GITHUB_TOKEN=[your token]
```

### create settings.yml

```sample.yml
rules:
  - name: master_protection
    branch_name: master
    protection_request:
      required_status_checks:
        strict: true
        contexts: []
      enforce_admins: true
      required_pull_request_reviews:
        dismiss_stale_reviews: false
        require_code_owner_reviews: true
        required_approving_review_count: 1
        dismissal_restrictions:
#          users: []
#          teams: []
      restrictions:
#        users: []
#        teams: []

owner: wasanx25
repos:
  - name: ghsts
    rules:
    - master_protection
```
