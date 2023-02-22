# GitHub Pull Request Aggregator

## About

- This command line tool is used to aggregate GitHub pull requests.

## Usage

### Configuration

```yaml
# config.yaml
elements:
  - source_owner: orgname0
    source_repo: reponame0
    base_branch: main
  - source_owner: orgname1
    source_repo: reponame1
    base_branch: develop
  - source_owner: orgname2
    source_repo: reponame2
    base_branch: main
```

### Build and Run

Execute the following command to run the tutorial:

```bash
make build
bin/tutorial -feature ghpra
```

Select the config to use:

```bash
bin/tutorial -feature ghpra
>> selected feature: ghpra
>> Start GitHub Pull Request Aggregation!
Select the config:
0 : orgname0 / reponame0 : branchname0
1 : orgname1 / reponame1 : branchname1
2 : orgname2 / reponame2 : branchname2

# input the number of the config
0

Authenticated user: username
****

PullRequestTitle0
PullRequestTitle1
PullRequestTitle2
...

Writing to a file: /path/to/file
```