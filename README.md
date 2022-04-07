# Use

## You can also use the [Demo Server](https://github.com/yihong0618/github-readme-stats-server) for test

- Update your README

Add a comment to your `README.md` like this:

```md
<!--START_SECTION:my_github-->
<!--END_SECTION:my_github-->
```

- Write your own `yml` file

[Sample](https://github.com/yihong0618/2021)

```yml
name: GitHub README STATS

on:
  workflow_dispatch:
  schedule:
    - cron: "0 0 * * *"
  push:
    branches:
      - main

env:
  GITHUB_NAME: yihong0618
  GITHUB_EMAIL: zouzou0208@gmail.com
  STARRED_NUM: 10

jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: My GitHub Status
        uses: yihong0618/github-readme-stats@main
        with:
          # if you also want to send TELE
          TELEGRAM_TOKEN: ${{ secrets.TELE_TOKEN }}
          TELEGRAM_CHAT_ID: ${{ secrets.TELE_CHAT_ID }}
          STARRED_NUM: ${{ env.STARRED_NUM }}

      - name: Push README
        uses: github-actions-x/commit@v2.6
        with:
          github-token: ${{ secrets.G_T }}
          # In this example, you can also use the ${{ secrets.GITHUB_TOKEN }} variable 
          # Permissions for the GITHUB_TOKEN : https://docs.github.com/en/free-pro-team@latest/actions/reference/authentication-in-a-workflow#permissions-for-the-github_token
        
          # If you need more precise Token permission control , you can create a personal access token and set it as a secret in your repository .
          commit-message: "Refresh README (GITHUB STATUS)"
          files: README.md
          rebase: "true"
          name: ${{ env.GITHUB_NAME }}
          email: ${{ env.GITHUB_EMAIL }}
```



# [My example](https://github.com/yihong0618/2021).

## My GitHub Status
<img align="middle" src="https://github-readme-stats-1.yihong0618.vercel.app/api?username=yihong0618&show_icons=true&&&hide_title=true" />

<!--START_SECTION:my_github-->
## The repos I created
| ID  |                                          REPO                                          |   START    |   UPDATE   |  LAUGUAGE  | STARS |
|-----|----------------------------------------------------------------------------------------|------------|------------|------------|-------|
|   1 | [running_page](https://github.com/yihong0618/running_page)                             | 2020-09-17 | 2021-12-09 | Python     |  1501 |
|   2 | [GitHubPoster](https://github.com/yihong0618/GitHubPoster)                             | 2021-04-21 | 2021-12-09 | Python     |   565 |
|   3 | [gitblog](https://github.com/yihong0618/gitblog)                                       | 2019-07-18 | 2021-12-09 | Python     |   326 |
|   4 | [2021](https://github.com/yihong0618/2021)                                             | 2020-12-21 | 2021-12-09 | Python     |   185 |
|   5 | [2020](https://github.com/yihong0618/2020)                                             | 2020-01-10 | 2021-12-03 | C          |   123 |
|   6 | [gaycore](https://github.com/yihong0618/gaycore)                                       | 2019-02-18 | 2021-11-20 | Python     |    90 |
|   7 | [iBeats](https://github.com/yihong0618/iBeats)                                         | 2021-06-11 | 2021-12-02 | Python     |    83 |
|   8 | [vscode-gcores](https://github.com/yihong0618/vscode-gcores)                           | 2020-01-04 | 2021-12-03 | TypeScript |    63 |
|   9 | [github-readme-stats](https://github.com/yihong0618/github-readme-stats)               | 2020-12-24 | 2021-12-09 | Go         |    63 |
|  10 | [dalian-IT](https://github.com/yihong0618/dalian-IT)                                   | 2021-04-07 | 2021-12-02 | md         |    52 |
|  11 | [shanbay_remember](https://github.com/yihong0618/shanbay_remember)                     | 2020-12-02 | 2021-06-24 | JavaScript |    35 |
|  12 | [duolingo_remember](https://github.com/yihong0618/duolingo_remember)                   | 2021-01-18 | 2021-09-11 | Python     |    35 |
|  13 | [nbnhhsh-cli](https://github.com/yihong0618/nbnhhsh-cli)                               | 2021-07-09 | 2021-11-12 | Python     |    29 |
|  14 | [pengdu_helper](https://github.com/yihong0618/pengdu_helper)                           | 2021-09-09 | 2021-12-09 | Go         |    24 |
|  15 | [gcores_calendar](https://github.com/yihong0618/gcores_calendar)                       | 2020-08-24 | 2021-12-09 | JavaScript |    21 |
|  16 | [running_skyline](https://github.com/yihong0618/running_skyline)                       | 2021-03-02 | 2021-11-14 | Python     |    19 |
|  17 | [my_kindle_stats](https://github.com/yihong0618/my_kindle_stats)                       | 2021-11-18 | 2021-11-25 | Python     |    14 |
|  18 | [blog](https://github.com/yihong0618/blog)                                             | 2020-06-22 | 2021-12-08 | JavaScript |    14 |
|  19 | [Runtastic](https://github.com/yihong0618/Runtastic)                                   | 2020-07-24 | 2021-07-30 | Python     |     8 |
|  20 | [Python365](https://github.com/yihong0618/Python365)                                   | 2019-09-05 | 2021-07-09 | Python     |     6 |
|  21 | [github-readme-stats-server](https://github.com/yihong0618/github-readme-stats-server) | 2021-12-08 | 2021-12-09 | HTML       |     5 |
|  22 | [yihong0618](https://github.com/yihong0618/yihong0618)                                 | 2020-07-16 | 2021-12-02 | md         |     4 |
|  23 | [run](https://github.com/yihong0618/run)                                               | 2021-08-16 | 2021-12-09 | Python     |     2 |
|  24 | [edocteel001](https://github.com/yihong0618/edocteel001)                               | 2019-11-12 | 2020-05-18 | JavaScript |     1 |
|  25 | [github_upstream_script](https://github.com/yihong0618/github_upstream_script)         | 2021-05-08 | 2021-05-08 | Python     |     1 |
|  26 | [gaycore-server](https://github.com/yihong0618/gaycore-server)                         | 2019-02-18 | 2020-11-02 | Go         |     0 |
|  27 | [test_svg](https://github.com/yihong0618/test_svg)                                     | 2021-03-18 | 2021-09-17 | md         |     0 |
| sum |                                                                                        |            |            |            |  3269 |

## The repos I contributed to
| ID  |                                         REPO                                         |                                   FIRSTDATE                                   |                                   LASTEDATE                                   |                                             PRCOUNT                                             |
|-----|--------------------------------------------------------------------------------------|-------------------------------------------------------------------------------|-------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
|   1 | [GpxTrackPoster](https://github.com/flopp/GpxTrackPoster)                            | [2019-08-06](https://github.com/flopp/GpxTrackPoster/pull/39)                 | [2021-03-20](https://github.com/flopp/GpxTrackPoster/pull/87)                 | [12](https://github.com/flopp/GpxTrackPoster/pulls?q=is%3Apr+author%3Ayihong0618)               |
|   2 | [leetcode-cli](https://github.com/leetcode-tools/leetcode-cli)                       | [2019-11-29](https://github.com/leetcode-tools/leetcode-cli/pull/31)          | [2020-08-21](https://github.com/leetcode-tools/leetcode-cli/pull/49)          | [8](https://github.com/leetcode-tools/leetcode-cli/pulls?q=is%3Apr+author%3Ayihong0618)         |
|   3 | [vscode-leetcode](https://github.com/LeetCode-OpenSource/vscode-leetcode)            | [2019-12-03](https://github.com/LeetCode-OpenSource/vscode-leetcode/pull/487) | [2020-07-22](https://github.com/LeetCode-OpenSource/vscode-leetcode/pull/602) | [6](https://github.com/LeetCode-OpenSource/vscode-leetcode/pulls?q=is%3Apr+author%3Ayihong0618) |
|   4 | [taichi](https://github.com/taichi-dev/taichi)                                       | [2021-09-23](https://github.com/taichi-dev/taichi/pull/2979)                  | [2021-10-23](https://github.com/taichi-dev/taichi/pull/3256)                  | [5](https://github.com/taichi-dev/taichi/pulls?q=is%3Apr+author%3Ayihong0618)                   |
|   5 | [nrc-exporter](https://github.com/yasoob/nrc-exporter)                               | [2020-07-05](https://github.com/yasoob/nrc-exporter/pull/2)                   | [2020-10-07](https://github.com/yasoob/nrc-exporter/pull/11)                  | [5](https://github.com/yasoob/nrc-exporter/pulls?q=is%3Apr+author%3Ayihong0618)                 |
|   6 | [awesome-cn-cafe](https://github.com/ElaWorkshop/awesome-cn-cafe)                    | [2020-08-04](https://github.com/ElaWorkshop/awesome-cn-cafe/pull/167)         | [2020-08-10](https://github.com/ElaWorkshop/awesome-cn-cafe/pull/170)         | [3](https://github.com/ElaWorkshop/awesome-cn-cafe/pulls?q=is%3Apr+author%3Ayihong0618)         |
|   7 | [kb](https://github.com/gnebbia/kb)                                                  | [2020-09-21](https://github.com/gnebbia/kb/pull/13)                           | [2020-09-23](https://github.com/gnebbia/kb/pull/28)                           | [3](https://github.com/gnebbia/kb/pulls?q=is%3Apr+author%3Ayihong0618)                          |
|   8 | [strava-datasource](https://github.com/grafana/strava-datasource)                    | [2021-04-13](https://github.com/grafana/strava-datasource/pull/34)            | [2021-05-13](https://github.com/grafana/strava-datasource/pull/39)            | [2](https://github.com/grafana/strava-datasource/pulls?q=is%3Apr+author%3Ayihong0618)           |
|   9 | [Tweet2Telegram](https://github.com/NeverBehave/Tweet2Telegram)                      | [2021-05-21](https://github.com/NeverBehave/Tweet2Telegram/pull/7)            | [2021-05-21](https://github.com/NeverBehave/Tweet2Telegram/pull/7)            | [2](https://github.com/NeverBehave/Tweet2Telegram/pulls?q=is%3Apr+author%3Ayihong0618)          |
|  10 | [nebula-python](https://github.com/vesoft-inc/nebula-python)                         | [2021-05-19](https://github.com/vesoft-inc/nebula-python/pull/106)            | [2021-05-20](https://github.com/vesoft-inc/nebula-python/pull/108)            | [2](https://github.com/vesoft-inc/nebula-python/pulls?q=is%3Apr+author%3Ayihong0618)            |
|  11 | [py-staticmaps](https://github.com/flopp/py-staticmaps)                              | [2020-09-20](https://github.com/flopp/py-staticmaps/pull/7)                   | [2021-03-24](https://github.com/flopp/py-staticmaps/pull/17)                  | [2](https://github.com/flopp/py-staticmaps/pulls?q=is%3Apr+author%3Ayihong0618)                 |
|  12 | [python-garminconnect](https://github.com/cyberjunky/python-garminconnect)           | [2021-02-26](https://github.com/cyberjunky/python-garminconnect/pull/43)      | [2021-05-25](https://github.com/cyberjunky/python-garminconnect/pull/49)      | [2](https://github.com/cyberjunky/python-garminconnect/pulls?q=is%3Apr+author%3Ayihong0618)     |
|  13 | [activities](https://github.com/flopp/activities)                                    | [2020-07-09](https://github.com/flopp/activities/pull/41)                     | [2020-07-14](https://github.com/flopp/activities/pull/44)                     | [2](https://github.com/flopp/activities/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|  14 | [iredis](https://github.com/laixintao/iredis)                                        | [2019-12-30](https://github.com/laixintao/iredis/pull/184)                    | [2020-09-16](https://github.com/laixintao/iredis/pull/360)                    | [2](https://github.com/laixintao/iredis/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|  15 | [notion-avatar](https://github.com/Mayandev/notion-avatar)                           | [2021-09-28](https://github.com/Mayandev/notion-avatar/pull/1)                | [2021-09-28](https://github.com/Mayandev/notion-avatar/pull/1)                | [1](https://github.com/Mayandev/notion-avatar/pulls?q=is%3Apr+author%3Ayihong0618)              |
|  16 | [LearnJapan](https://github.com/wizicer/LearnJapan)                                  | [2020-03-31](https://github.com/wizicer/LearnJapan/pull/2)                    | [2020-03-31](https://github.com/wizicer/LearnJapan/pull/2)                    | [1](https://github.com/wizicer/LearnJapan/pulls?q=is%3Apr+author%3Ayihong0618)                  |
|  17 | [awesome-cn-cafe-web](https://github.com/antfu/awesome-cn-cafe-web)                  | [2020-08-18](https://github.com/antfu/awesome-cn-cafe-web/pull/5)             | [2020-08-18](https://github.com/antfu/awesome-cn-cafe-web/pull/5)             | [1](https://github.com/antfu/awesome-cn-cafe-web/pulls?q=is%3Apr+author%3Ayihong0618)           |
|  18 | [tokei-pie](https://github.com/laixintao/tokei-pie)                                  | [2021-11-19](https://github.com/laixintao/tokei-pie/pull/2)                   | [2021-11-19](https://github.com/laixintao/tokei-pie/pull/2)                   | [1](https://github.com/laixintao/tokei-pie/pulls?q=is%3Apr+author%3Ayihong0618)                 |
|  19 | [stravalib](https://github.com/hozn/stravalib)                                       | [2021-08-18](https://github.com/hozn/stravalib/pull/218)                      | [2021-08-18](https://github.com/hozn/stravalib/pull/218)                      | [1](https://github.com/hozn/stravalib/pulls?q=is%3Apr+author%3Ayihong0618)                      |
|  20 | [olo](https://github.com/yetone/olo)                                                 | [2021-04-12](https://github.com/yetone/olo/pull/91)                           | [2021-04-12](https://github.com/yetone/olo/pull/91)                           | [1](https://github.com/yetone/olo/pulls?q=is%3Apr+author%3Ayihong0618)                          |
|  21 | [awesome-database-learning](https://github.com/pingcap/awesome-database-learning)    | [2021-05-11](https://github.com/pingcap/awesome-database-learning/pull/37)    | [2021-05-11](https://github.com/pingcap/awesome-database-learning/pull/37)    | [1](https://github.com/pingcap/awesome-database-learning/pulls?q=is%3Apr+author%3Ayihong0618)   |
|  22 | [help-to-be-helped](https://github.com/xiaolai/help-to-be-helped)                    | [2020-02-04](https://github.com/xiaolai/help-to-be-helped/pull/4)             | [2020-02-04](https://github.com/xiaolai/help-to-be-helped/pull/4)             | [1](https://github.com/xiaolai/help-to-be-helped/pulls?q=is%3Apr+author%3Ayihong0618)           |
|  23 | [GadioVideo](https://github.com/rabbitism/GadioVideo)                                | [2019-09-25](https://github.com/rabbitism/GadioVideo/pull/16)                 | [2019-09-25](https://github.com/rabbitism/GadioVideo/pull/16)                 | [1](https://github.com/rabbitism/GadioVideo/pulls?q=is%3Apr+author%3Ayihong0618)                |
|  24 | [highlight](https://github.com/wenyan-lang/highlight)                                | [2020-09-08](https://github.com/wenyan-lang/highlight/pull/4)                 | [2020-09-08](https://github.com/wenyan-lang/highlight/pull/4)                 | [1](https://github.com/wenyan-lang/highlight/pulls?q=is%3Apr+author%3Ayihong0618)               |
|  25 | [hub-mirror-action](https://github.com/Yikun/hub-mirror-action)                      | [2021-04-09](https://github.com/Yikun/hub-mirror-action/pull/101)             | [2021-04-09](https://github.com/Yikun/hub-mirror-action/pull/101)             | [1](https://github.com/Yikun/hub-mirror-action/pulls?q=is%3Apr+author%3Ayihong0618)             |
|  26 | [TopList](https://github.com/tophubs/TopList)                                        | [2019-08-19](https://github.com/tophubs/TopList/pull/13)                      | [2019-08-19](https://github.com/tophubs/TopList/pull/13)                      | [1](https://github.com/tophubs/TopList/pulls?q=is%3Apr+author%3Ayihong0618)                     |
|  27 | [gitlab-skyline](https://github.com/felixgomez/gitlab-skyline)                       | [2021-03-02](https://github.com/felixgomez/gitlab-skyline/pull/6)             | [2021-03-02](https://github.com/felixgomez/gitlab-skyline/pull/6)             | [1](https://github.com/felixgomez/gitlab-skyline/pulls?q=is%3Apr+author%3Ayihong0618)           |
|  28 | [MangaLineExtraction_PyTorch](https://github.com/ljsabc/MangaLineExtraction_PyTorch) | [2021-09-22](https://github.com/ljsabc/MangaLineExtraction_PyTorch/pull/3)    | [2021-09-22](https://github.com/ljsabc/MangaLineExtraction_PyTorch/pull/3)    | [1](https://github.com/ljsabc/MangaLineExtraction_PyTorch/pulls?q=is%3Apr+author%3Ayihong0618)  |
|  29 | [nebula](https://github.com/vesoft-inc/nebula)                                       | [2021-05-17](https://github.com/vesoft-inc/nebula/pull/2476)                  | [2021-05-17](https://github.com/vesoft-inc/nebula/pull/2476)                  | [1](https://github.com/vesoft-inc/nebula/pulls?q=is%3Apr+author%3Ayihong0618)                   |
|  30 | [build-your-own-vue](https://github.com/jackiewillen/build-your-own-vue)             | [2020-01-16](https://github.com/jackiewillen/build-your-own-vue/pull/1)       | [2020-01-16](https://github.com/jackiewillen/build-your-own-vue/pull/1)       | [1](https://github.com/jackiewillen/build-your-own-vue/pulls?q=is%3Apr+author%3Ayihong0618)     |
|  31 | [xrkffgg](https://github.com/xrkffgg/xrkffgg)                                        | [2021-03-18](https://github.com/xrkffgg/xrkffgg/pull/3)                       | [2021-03-18](https://github.com/xrkffgg/xrkffgg/pull/3)                       | [1](https://github.com/xrkffgg/xrkffgg/pulls?q=is%3Apr+author%3Ayihong0618)                     |
| sum |                                                                                      |                                                                               |                                                                               |                                                                                              73 |

## The repos I stared (random 10)
| ID |                                             REPO                                             | STAREDDATE |  LAUGUAGE  | LATESTUPDATE |
|----|----------------------------------------------------------------------------------------------|------------|------------|--------------|
|  1 | [PaddleGAN](https://github.com/PaddlePaddle/PaddleGAN)                                       | 2021-03-04 | Python     | 2021-12-10   |
|  2 | [taxi-demo](https://github.com/kevinqqnj/taxi-demo)                                          | 2019-11-07 | Python     | 2019-11-07   |
|  3 | [blog](https://github.com/gwuhaolin/blog)                                                    | 2018-02-05 | EJS        | 2021-12-09   |
|  4 | [go-fundamental-programming](https://github.com/unknwon/go-fundamental-programming)          | 2018-05-21 | Go         | 2021-12-09   |
|  5 | [machine_learning_examples](https://github.com/lazyprogrammer/machine_learning_examples)     | 2020-06-08 | Python     | 2021-12-09   |
|  6 | [NeteaseCloudMusic-Now-Playing](https://github.com/SumiMakito/NeteaseCloudMusic-Now-Playing) | 2020-08-04 | C          | 2021-01-30   |
|  7 | [OneForAll](https://github.com/shmilylty/OneForAll)                                          | 2019-08-07 | Python     | 2021-12-09   |
|  8 | [polyline](https://github.com/frederickjansen/polyline)                                      | 2020-09-17 | Python     | 2021-11-09   |
|  9 | [daily-anime](https://github.com/deepred5/daily-anime)                                       | 2019-12-28 | JavaScript | 2021-11-11   |
| 10 | [tornado-websocket-example](https://github.com/hiroakis/tornado-websocket-example)           | 2018-11-28 | HTML       | 2021-11-29   |


<!--END_SECTION:my_github-->
