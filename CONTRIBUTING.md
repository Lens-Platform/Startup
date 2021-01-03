# Contributing

--

First, thank you for contributing to the Blackspace Platform! The goal of this document is to
provide everything you need to start contributing to Blackspace properly. The
following TOC is sorted progressively, starting with the basics and
expanding into more specifics.

<!-- MarkdownTOC autolink="true" style="ordered" indent="   " -->

- [Contributing](#contributing)
  - [Assumptions](#assumptions)
  - [Your First Contribution](#your-first-contribution)
  - [Change Control](#change-control)
  - [Onboarding](#onboarding)
    - [Git Branches](#git-branches)
    - [Git Commits](#git-commits)
      - [Style](#style)
      - [Signing-off](#signing-off)
    - [Github Pull Requests](#github-pull-requests)
      - [Title](#title)
      - [Reviews & Approvals](#reviews--approvals)
      - [Merge Style](#merge-style)
    - [CI](#ci)
  - [Development](#development)
    - [Setup](#setup)
      - [Using a Docker environment](#using-a-docker-environment)
    - [The Basics](#the-basics)
      - [Directory Structure](#directory-structure)
      - [Makefile](#makefile)
      - [Code Style](#code-style)
      - [Feature flags](#feature-flags)
      - [Dependencies](#dependencies)
    - [Guidelines](#guidelines)
      - [Service Healthchecks](#service-healthchecks)
    - [Testing](#testing)
      - [Sample Logs](#sample-logs)
      - [Tips and Tricks](#tips-and-tricks)
    - [Benchmarking](#benchmarking)
    - [Profiling](#profiling)
    - [Kubernetes](#kubernetes)
      - [Kubernetes Dev Flow](#kubernetes-dev-flow)
        - [Requirements](#requirements)
        - [The dev flow](#the-dev-flow)
        - [Troubleshooting](#troubleshooting)
      - [Kubernetes E2E tests](#kubernetes-e2e-tests)
        - [Requirements](#requirements-1)
        - [Running the E2E tests](#running-the-e2e-tests)
  - [Humans](#humans)
    - [Documentation](#documentation)
    - [Changelog](#changelog)
      - [What makes a highlight noteworthy?](#what-makes-a-highlight-noteworthy)
      - [How is a highlight different from a blog post?](#how-is-a-highlight-different-from-a-blog-post)
  - [Security](#security)
  - [Legal](#legal)
    - [DCO](#dco)
      - [Trivial changes](#trivial-changes)
  - [FAQ](#faq)
    - [What if I forgot to sign my commits?](#what-if-i-forgot-to-sign-my-commits)

<!-- /MarkdownTOC -->

## Assumptions

1. **You're familiar with [Github](https://github.com) and the pull request
   workflow.**
2. **You've read Blackspace's [docs](https://github.io/blackspace/docs/).**
3. **You know about the [Blackspace community](https://gitter.im/BlackspaceInc/community).
   Please use this for help.**
4. You are familiar with and have read the **style guides** for this project present in the __/.style_guide/*__

## Your First Contribution
Run the following scripts/commands in the terminal to download conform and set up a commit message hook locally.
```bash
go get github.com/talos-systems/conform
```

```bash
cat <<EOF | tee .git/hooks/commit-msg
#!/bin/sh

conform enforce --commit-msg-file \$1
EOF
chmod +x .git/hooks/commit-msg
```
0. Read the style guides in __/.style_guide/*__
1. Ensure your change has an issue! Find an
   [existing issue][urls.existing_issues] or [open a new issue][urls.new_issue].
   - This is where you can get a feel if the change will be accepted or not.
     Changes that are questionable will have a `needs: approval` label.
2. Once an issue is obtained, [fork the Blackspace repository from dev branch][urls.fork_repo] in your own local instance.
3. [Create a new Git branch off of the dev branch][urls.create_branch].
   ```bash
      git clone https://github.com/Lens-Platform/Platform.git
      git checkout -b new_branch_name dev
   ```
4. Review the Blackspace [workflow](#workflow) and [development](#development) guidelines.
5. Make your changes.
6. [Submit the branch as a pull request][urls.submit_pr] to the main Blackspace
   repo `dev branch`. A Blackspace team member should comment and/or review your pull request
   within a few days. Although, depending on the circumstances, it may take
   longer.

## Change Control

## Onboarding

_Start_ documentation site locally to read onboarding resources.
```shell
make doc-site
```

### Git Branches

_All_ changes must be made in a branch and submitted as [pull requests](#pull-requests).
Blackspace does not adopt any type of branch naming style, but please use something
descriptive of your changes.

### Git Commits

_All_ commits should be concise and adequately describe the change performed. Please ensure you maintain small commits
and perform on thing at a time. Each commit should be associated with one distinct change. Commit should follow a distinct framework
> Commits should have follow the following framework
>
> git commit -sm "[Tag]:[Description]"
>
> Acceptable Tags :
>
>> fix:, feat:, build:, chore:, ci:, docs:, style:, refactor:, perf:, test:
>
> example: git commit -sm "chore: remove unecessary configs"

#### Style

Please ensure your commits are small and focused; they should tell a story of
your change. This helps reviewers to follow your changes, especially for more
complex changes.

#### Signing-off

Your commits must include a [DCO](https://developercertificate.org/) signature.
This is simpler than it sounds; it just means that all of your commits
must contain:

```text
Signed-off-by: Joe Smith <joe.smith@email.com>
```

Git makes this easy by adding the `-s` or `--signoff` flags when you commit:

```bash
git commit -sm 'My commit message'
```

We also included a `make signoff` target that handles this for you if
you forget.

### Github Pull Requests

Once your changes are ready you must submit your branch as a [pull \
request](https://github.com/timberio/Blackspace/pulls).

#### Title

The pull request title must follow the format outlined in the [conventional \
commits spec](https://www.conventionalcommits.org).
[Conventional commits](https://www.conventionalcommits.org) is a standardized
format for commit messages. Blackspace only requires this format for commits on
the `master` branch. And because Blackspace squashes commits before merging
branches, this means that only the pull request title must conform to this
format. Blackspace performs a pull request check to verify the pull request title
in case you forget.

A list of allowed sub-categories is defined
[here](https://github.com/Lens-Platform/Platform/tree/master/.github).

The following are all good examples of pull request titles:

```text
feature(new sink): new `xyz` sink
feature(tcp source): add foo bar baz feature
fix(tcp source): fix foo bar baz bug
chore: improve build process
docs: fix typos
```

#### Reviews & Approvals

All pull requests must be reviewed and approved by at least one Blackspace team
member. The review process is outlined in the [Review guide](REVIEWING.md).

#### Merge Style

All pull requests are squashed and merged. We generally discourage large pull
requests that are over 300-500 lines of diff. If you would like to propose
a change that is larger we suggest coming onto the team's gitter or slack channel and
discuss it with one of your counterparts. This way we can talk through the
solution and discuss if a change that large is even needed! This overall
will produce a quicker response to the change and likely produce code that
aligns better with our process.

### CI

Currently Blackspace uses Github Actions to run tests. The workflows are defined in
`.github/workflows`.

Tests are run for all changes except those that have the label:

```text
ci-condition: skip
```

Github Actions is responsible for releasing updated versions of Blackspace through
various channels.

Some long running tests will only run daily, rather than on every pull request.
If needed, an administrator can kick off these tests manually.

## Development

### Setup

We're super excited to have you interested in working on Blackspace! Before you start you should pick how you want to develop.

For small or first-time contributions, we recommend the Docker method. If you do a lot of contributing, try adopting the Nix method! It'll be way faster and feel more smooth. Prefer to do it yourself? That's fine too!

#### Using a Docker environment

> **Targets:** You can use this method to produce AARCH64, Arm6/7, as well as x86/64 Linux builds.

Since not everyone has a full working native environment, or can use Nix, we will take our Nix environment and stuff it into orchestrated Docker containers!

This is ideal for users who want it to "Just work" and just want to start contributing. It's also what we use for our CI, so you know if it breaks we can't do anything else until we fix it. 😉

**Before you go farther, install Docker through your official package manager, or from the [Docker](https://docs.docker.com/get-docker/) sites.**

By default, `make environment` style tasks will do a `docker pull` from Github's container repository, you can **optionally** build your own environment while you make your morning coffee ☕:

```bash
# Optional: Only if you want to go make a coffee
# sets up environment in which services will run (development or production)
```

Now that you have your coffee, you can enter the shell!

```bash
# Spins up services locally
# In Debug Mode


# In Production mode
```

Now you can use the jobs detailed in **"Bring your own toolbox"** below.

### The Basics

#### Directory Structure

- [`/.github`](/.github) - Store CI workflows for Blackspace.
- [`/.meta`](/.meta) - Project metadata used to generate documentation.
- [`/.style_guide`](/.style_guide) -  Witholds project styleguides and best practices
- [`/benches`](/benches) - Internal benchmarks.
- [`/docs`](/docs) -  Store documentation.
- [`/istio-manifests`](/istio-manifests) -  Istio service mesh and production manifests.
- [`/kubernetes-manifests`](/kubernetes-manifests) -  Kubernetes production manifests.
- [`/linkerd-manifests`](/linkerd-manifests) -  Linkerd service mesh and production manifests.
- [`/monitoring`](/monitoring) -  Production and development configurations for monitoring purposes.
- [`/pb`](/pb) -  Protobuf schema definitions common across all backend services.
- [`/release`](/release) -  Release artifacts.
- [`/scripts`](/scripts) - Scripts used to generate docs and maintain the repo.
- [`/src`](/src) - Blackspace source.
- [`/distribution`](/distribution) - Distribution artifacts for various targets.
- [`/test-harness`](/test-harness) - Utility for E2E tests.

#### Makefile

Blackspace includes a [`Makefile`](/Makefile) in the root of the repo. This serves
as a high-level interface for common commands. Running `make` will produce
a list of make targets with descriptions. These targets will be referenced
throughout this document.

#### Code Style

We use `gofmt` to format our code and CI will verify that your
code follows
this format style. To run the following command make sure `gofmt` has been
installed on the stable toolchain locally.

```bash
# To format the code
make fmt
```

#### Feature flags

When a new feature or microservice is added, it has to be put
behind a feature flag with the corresponding name. This ensures that it is
possible to customize Blackspace builds as well as turn it on or off in production especially if we foresee issues arising.

In addition, during chaos testing it is useful to
disable certain services entirely at runtime to better gain visibility over the system's performance in such scenarios.

#### Dependencies

Dependencies should be _carefully_ selected and avoided if possible. You can
see how dependencies are reviewed in the
[Reviewing guide](/REVIEWING.md#dependencies).

### Guidelines

#### Service Healthchecks

Services must implement a health check as a means for validating their configuration
against the environment and external systems. Ideally, this allows the system to
inform users of problems such as insufficient credentials, unreachable
endpoints, non-existent tables, etc. They're not perfect, however, since it's
impossible to exhaustively check for issues that may happen at runtime.

When implementing health checks, we prefer false positives to false negatives.
This means we would prefer that a health check pass and the service then fail than
to have the health check fail when the service would have been able to run
successfully.

A common cause of false negatives in health checks is performing an operation
that the service itself does not need. For example, listing all of the available S3
buckets and checking that the configured bucket is on that list. The S3 sink
doesn't need the ability to list all buckets, and a user that knows that may not
have permitted it to do so. In that case, the health check will fail due
to bad credentials even through its credentials are sufficient for normal
operation.

This leads to a general strategy of mimicking what the service itself does.
Unfortunately, the fact that health checks don't have real events available to
them leads to some limitations here. The most obvious example of this is with
services where the exact target of a write depends on the value of some field in
the event. It also pops up for services
where incoming events are expected to conform to a specific schema. In both
cases, random test data is reasonably likely to trigger a potentially
false-negative result. Even in simpler cases, we need to think about the effects
of writing test data and whether the user would find that surprising or
invasive. The answer usually depends on the system we're interfacing with.

In some cases, the right thing to do might be
nothing at all. If we require dynamic information to figure out what entity
that we're even dealing with, odds are very
low that we'll be able to come up with a way to meaningfully validate that it's
in working order. It's perfectly valid to have a health check that falls back to
doing nothing when there is a data dependency like this.

With all that in mind, here is a simple checklist to go over when writing a new
health check:

- [ ] Does this check perform different fallible operations from the service itself?
- [ ] Does this check have side effects the user would consider undesirable (e.g. data pollution)?
- [ ] Are there situations where this check would fail but the service would operate normally?

Not all of the answers need to be a hard "no", but we should think about the
likelihood that any "yes" would lead to false negatives and balance that against
the usefulness of the check as a whole for finding problems. Because we have the
option to disable individual health checks, there's an escape hatch for users
that fall into a false negative circumstance. Our goal should be to minimize the
likelihood of users needing to pull that lever while still making a good effort
to detect common problems.

### Testing

You can run Blackspace's tests via the `make test` command. Our tests either use Docker
compose to spin up mock services for testing or spins up the necessary service containers in a local minikube cluster.

#### Sample Logs

We use `flog` to build a sample set of log files to test sending logs from a
file to our local Kibana stack. This can be done with the following commands on mac with homebrew.
Installation instruction for flog can be found
[here](https://github.com/mingrammer/flog#installation).

```bash
flog --bytes $((100 * 1024 * 1024)) > sample.log
```

This will create a `100MiB` sample log file in the `sample.log` file.

#### Tips and Tricks

If you are developing a particular component and want to quickly iterate on unit
tests related only to this component, run the unit tests from the folder in which the component was developed.

### Benchmarking

All benchmarks are placed in the [`/benches`](/benches) folder. You can
run benchmarks via the `make benchmarks` command. In addition, Blackspace
maintains a full [test hardness][urls.Blackspace_test_harness] for complex
end-to-end integration and performance testing.

### Profiling

If you're trying to improve Blackspace's performance (or understand why your change
made it worse), profiling is a useful tool for seeing where time is being spent.

While there are a bunch of useful profiling tools, a simple place to get started
is with Linux's `perf`. Before getting started, you'll likely need to give
yourself access to collect stats:

```sh
echo -1 | sudo tee /proc/sys/kernel/perf_event_paranoid
```

You'll also want to edit `.env` files specific to the services you have updated and make sure that Blackspace is being built
with debug symbols in release mode. This ensures that you'll get human-readable
info in the eventual output:

```.env
debug = true
```

Then you can start up a release build of Blackspace with whatever config you're
interested in profiling.

```sh
make release
```

Once it's started, use the `ps` tool (or equivalent) to make a note of its PID.
We'll use this to tell `perf` which process we would like it to collect data
about.

The next step is somewhat dependent on the config you're testing. For this
example, let's assume you're using a simple TCP-mode socket source listening on
port 9000. Let's also assume that you have a large file of example input in
`access.log` (you can use a tool like `flog` to generate this).

With all that prepared, we can send our test input to Blackspace and collect data
while it is under load:

```sh
perf record -F99 --call-graph dwarf -p $Blackspace_PID socat -dd OPEN:access.log TCP:localhost:9000
```

This instructs `perf` to collect data from our already-running Blackspace process
for the duration of the `socat` command. The `-F` argument is the frequency at
which `perf` should sample the Blackspace call stack. Higher frequencies will
collect more data and produce more detailed output, but can produce enormous
amounts of data that take a very long time to process. Using `-F99` works well
when your input data is large enough to take a minute or more to process, but
feel free to adjust both input size and sampling frequency for your setup.

It's worth noting that this is not the normal way to profile programs with
`perf`. Usually you would simply run something like `perf record my_program` and
not have to worry about PIDs and such. We differ from this because we're only
interested in data about what Blackspace is doing while under load. Running it
directly under `perf` would collect data for the entire lifetime of the process,
including startup, shutdown, and idle time. By telling `perf` to collect data
only while the load generation command is running we get a more focused dataset
and don't have to worry about timing different commands in quick succession.

You'll now find a `perf.data` file in your current directory with all of the
information that was collected. There are different ways to process this, but
one of the most useful is to create
a [flamegraph](http://www.brendangregg.com/flamegraphs.html). For this we can
use the `inferno` tool (available via `cargo install`):

```sh
perf script | inferno-collapse-perf > stacks.folded
cat stacks.folded | inferno-flamegraph > flamegraph.svg
```

And that's it! You now have a flamegraph SVG file that can be opened and
navigated in your favorite web browser.

### Kubernetes

#### Kubernetes Dev Flow

There is a special flow for when you develop Blackspace especially since it is
designed to work with both docker compose and kubernetes.

This flow facilitates building Blackspace and deploying it into a cluster.

##### Requirements

There are some extra requirements besides what you'd normally need to work on
Blackspace:

- `linux` or `windows` system
- [`skaffold`](https://skaffold.dev/)
- [`docker`](https://www.docker.com/)
- [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [`kustomize`](https://kustomize.io/)
- [`minikube`](https://minikube.sigs.k8s.io/)-powered or other k8s cluster

##### The dev flow

Once you have the requirements, use the `make spin-up-kube` command.

That's it, just one command should take care of everything!

It will

1. build the various `Blackspace` service binaries in development mode,
2. build docker images from these binaries,
3. deploy `Blackspace services` into the Kubernetes cluster at your current kubectl context
   using the built docker image and a mix of our production deployment
   configuration defined in the directory of each service

As the result of invoking the `make spin-up-kube` command, you should see
a `kubernetes` process running on your local machine, printing the logs from the
various deployed `Blackspace service` instances.

To stop the process, press `Ctrl+C`, and wait for `kubernetes` to clean up
the cluster state and exit.

Additionally, you can configure [skaffold](https://skaffold.dev/docs/) to automatically set up the kubernetes env requirements
automatically so we can better focus on development.

##### Troubleshooting

#### Kubernetes E2E tests

Kubernetes integration has a lot of parts that can go wrong.

To cope with the complexity and ensure we maintain high quality, we use
E2E (end-to-end) tests.

> E2E tests normally run at CI, so there's typically no need to run them
> manually.

##### Requirements

- `kubernetes` cluster (`minikube` has special support, but any cluster should
  work)
- `docker`
- `kubectl`
- `bash`

Blackspace release artifacts are prepared for E2E tests, so the ability to do that
is required too, see Blackspace [docs](https://Blackspace.dev) for more details.

Also:

> Note: `minikube` has troubles running on ZFS systems. If you're using ZFS, we
> suggest using a cloud cluster or [`minik8s`](https://microk8s.io/) with local
> registry.

##### Running the E2E tests

To run the E2E tests, use the following command:

```shell
make run-tests "e2e"
```

## Humans

After making your change, you'll want to prepare it for Blackspace's users
(mostly humans). This usually entails updating documentation and announcing
your feature.

### Documentation

Documentation is very important to the Blackspace project! In order to keep things
simple for contributors, all reference documentation is derived from metadata
in the [`/.meta` directory](/.meta). If you add a service, change options,
or otherwise change anything user facing, you should update the relevant files
in the `/.meta` directory.

To ensure your change is valid, you can run `make check-meta`, which validates
your changes against the local `/.meta/.schema.json` file.

The actual website and documentation are generated on the
[`Blackspace-website`](https://github.com/Lens-Platform/Platform) repo.

### Changelog

Developers do not need to maintain the [`Changelog`](/CHANGELOG.md). This is
automatically generated via the `make release` command. This is made possible
by the use of [conventional commit](#title) titles.

#### What makes a highlight noteworthy?

It should offer meaningful value to users. This is inherently subjective and
it is impossible to define exact rules for this distinction. But we should be
cautious not to dillute the meaning of a highlight by producing low values
highlights.

#### How is a highlight different from a blog post?

Highlights are not blog posts. They are short one, maybe two, paragraph
announcements. Highlights should allude to, or link to, a blog post if
relevant.

Highlights alludes to an
upcoming blog post. This allows us to communicate a high-value performance
improvement without being blocked by an in-depth blog post.

## Security

Please see the [`SECURITY.md` file](/SECURITY.md).

## Legal

To protect all users of Blackspace, the following legal requirements are made.

### DCO

Blackspace requires all contributors to agree to the DCO. DCO stands for Developer
Certificate of Origin and is maintained by the
[Linux Foundation](https://www.linuxfoundation.org). It is an attestation
attached to every commit made by every developer. It ensures that all committed
code adheres to the [Blackspace license](LICENSE.md) (Apache 2.0).

#### Trivial changes

Trivial changes, such as spelling fixes, do not need to be signed.
patent licenses as well. These rights are granted to everyone.

## FAQ

### What if I forgot to sign my commits?

No probs! We made this simple with the [`signoff` Makefile target](Makefile):

```bash
make signoff
```

If you prefer to do this manually:

https://stackoverflow.com/questions/13043357/git-sign-off-previous-commits

[urls.aws_announcements]: https://aws.amazon.com/new/?whats-new-content-all.sort-by=item.additionalFields.postDateTime&whats-new-content-all.sort-order=desc&wn-featured-announcements.sort-by=item.additionalFields.numericSort&wn-featured-announcements.sort-order=asc
[urls.create_branch]: https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-and-deleting-branches-within-your-repository
[urls.existing_issues]: https://github.com/Lens-Platform/Platform/issues
[urls.fork_repo]: https://help.github.com/en/github/getting-started-with-github/fork-a-repo
[urls.github_sign_commits]: https://help.github.com/en/github/authenticating-to-github/signing-commits
[urls.new_issue]: https://github.com/Lens-Platform/Platform/issues/new
[urls.push_it_to_the_limit]: https://www.youtube.com/watch?v=ueRzA9GUj9c
[urls.performance_highlight]: https://github.com/Lens-Platform/Platform/highlights/2020-04-11-overall-performance-increase
[urls.submit_pr]: https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request-from-a-fork
[urls.Blackspace_test_harness]: https://github.com/timberio/vector-test-harness
