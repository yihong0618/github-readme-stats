# Use

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
  STARED_NUMBER: 10

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
          STARED_NUMBER: ${{ env.STARED_NUMBER }}

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
| ID  |                                    REPO                                    |   START    |   UPDATE   |  LAUGUAGE  | STARS |
|-----|----------------------------------------------------------------------------|------------|------------|------------|-------|
|   1 | [running_page](https://github.com/yihong0618/running_page)                 | 2020-09-17 | 2021-04-15 | Python     |  1185 |
|   2 | [2020](https://github.com/yihong0618/2020)                                 | 2020-01-10 | 2021-04-14 | C          |   112 |
|   3 | [gitblog](https://github.com/yihong0618/gitblog)                           | 2019-07-18 | 2021-04-14 | Python     |   106 |
|   4 | [gaycore](https://github.com/yihong0618/gaycore)                           | 2019-02-18 | 2021-01-24 | Python     |    88 |
|   5 | [vscode-gcores](https://github.com/yihong0618/vscode-gcores)               | 2020-01-04 | 2021-04-06 | TypeScript |    52 |
|   6 | [shanbay_remember](https://github.com/yihong0618/shanbay_remember)         | 2020-12-02 | 2021-03-28 | JavaScript |    35 |
|   7 | [github-readme-stats](https://github.com/yihong0618/github-readme-stats)   | 2020-12-24 | 2021-04-14 | Go         |    35 |
|   8 | [gcores_calendar](https://github.com/yihong0618/gcores_calendar)           | 2020-08-24 | 2021-04-15 | JavaScript |    21 |
|   9 | [2021](https://github.com/yihong0618/2021)                                 | 2020-12-21 | 2021-04-15 | Python     |    18 |
|  10 | [dalian-IT](https://github.com/yihong0618/dalian-IT)                       | 2021-04-07 | 2021-04-15 | md         |    16 |
|  11 | [duolingo_remember](https://github.com/yihong0618/duolingo_remember)       | 2021-01-18 | 2021-02-21 | Python     |    15 |
|  12 | [running_skyline](https://github.com/yihong0618/running_skyline)           | 2021-03-02 | 2021-03-10 | Python     |    14 |
|  13 | [blog](https://github.com/yihong0618/blog)                                 | 2020-06-22 | 2021-04-14 | JavaScript |    10 |
|  14 | [Runtastic](https://github.com/yihong0618/Runtastic)                       | 2020-07-24 | 2020-11-16 | Python     |     7 |
|  15 | [Python365](https://github.com/yihong0618/Python365)                       | 2019-09-05 | 2021-02-23 | Python     |     4 |
|  16 | [yihong0618](https://github.com/yihong0618/yihong0618)                     | 2020-07-16 | 2021-01-20 | md         |     1 |
|  17 | [edocteel001](https://github.com/yihong0618/edocteel001)                   | 2019-11-12 | 2020-05-18 | JavaScript |     1 |
|  18 | [collections](https://github.com/yihong0618/collections)                   | 2019-03-05 | 2019-03-06 | md         |     0 |
|  19 | [hoingyi_bot](https://github.com/yihong0618/hoingyi_bot)                   | 2019-10-16 | 2020-05-25 | Shell      |     0 |
|  20 | [test_svg](https://github.com/yihong0618/test_svg)                         | 2021-03-18 | 2021-03-20 | md         |     0 |
|  21 | [gaycore-server](https://github.com/yihong0618/gaycore-server)             | 2019-02-18 | 2020-11-02 | Go         |     0 |
|  22 | [Frontend100](https://github.com/yihong0618/Frontend100)                   | 2019-10-31 | 2019-12-05 | HTML       |     0 |
|  23 | [zouzou0208.github.io](https://github.com/yihong0618/zouzou0208.github.io) | 2018-03-06 | 2018-03-06 | HTML       |     0 |
| sum |                                                                            |            |            |            |  1720 |

## The repos I contributed to
| ID  |                                    REPO                                    |                                   FIRSTDATE                                   |                                   LASTEDATE                                   |                                             PRCOUNT                                             |
|-----|----------------------------------------------------------------------------|-------------------------------------------------------------------------------|-------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
|   1 | [GpxTrackPoster](https://github.com/flopp/GpxTrackPoster)                  | [2019-08-06](https://github.com/flopp/GpxTrackPoster/pull/39)                 | [2021-03-20](https://github.com/flopp/GpxTrackPoster/pull/87)                 | [12](https://github.com/flopp/GpxTrackPoster/pulls?q=is%3Apr+author%3Ayihong0618)               |
|   2 | [leetcode-cli](https://github.com/leetcode-tools/leetcode-cli)             | [2019-11-29](https://github.com/leetcode-tools/leetcode-cli/pull/31)          | [2020-08-21](https://github.com/leetcode-tools/leetcode-cli/pull/49)          | [8](https://github.com/leetcode-tools/leetcode-cli/pulls?q=is%3Apr+author%3Ayihong0618)         |
|   3 | [vscode-leetcode](https://github.com/LeetCode-OpenSource/vscode-leetcode)  | [2019-12-03](https://github.com/LeetCode-OpenSource/vscode-leetcode/pull/487) | [2020-07-22](https://github.com/LeetCode-OpenSource/vscode-leetcode/pull/602) | [6](https://github.com/LeetCode-OpenSource/vscode-leetcode/pulls?q=is%3Apr+author%3Ayihong0618) |
|   4 | [nrc-exporter](https://github.com/yasoob/nrc-exporter)                     | [2020-07-05](https://github.com/yasoob/nrc-exporter/pull/2)                   | [2020-10-07](https://github.com/yasoob/nrc-exporter/pull/11)                  | [5](https://github.com/yasoob/nrc-exporter/pulls?q=is%3Apr+author%3Ayihong0618)                 |
|   5 | [kb](https://github.com/gnebbia/kb)                                        | [2020-09-21](https://github.com/gnebbia/kb/pull/13)                           | [2020-09-23](https://github.com/gnebbia/kb/pull/28)                           | [3](https://github.com/gnebbia/kb/pulls?q=is%3Apr+author%3Ayihong0618)                          |
|   6 | [awesome-cn-cafe](https://github.com/ElaWorkshop/awesome-cn-cafe)          | [2020-08-04](https://github.com/ElaWorkshop/awesome-cn-cafe/pull/167)         | [2020-08-10](https://github.com/ElaWorkshop/awesome-cn-cafe/pull/170)         | [3](https://github.com/ElaWorkshop/awesome-cn-cafe/pulls?q=is%3Apr+author%3Ayihong0618)         |
|   7 | [activities](https://github.com/flopp/activities)                          | [2020-07-09](https://github.com/flopp/activities/pull/41)                     | [2020-07-14](https://github.com/flopp/activities/pull/44)                     | [2](https://github.com/flopp/activities/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|   8 | [iredis](https://github.com/laixintao/iredis)                              | [2019-12-30](https://github.com/laixintao/iredis/pull/184)                    | [2020-09-16](https://github.com/laixintao/iredis/pull/360)                    | [2](https://github.com/laixintao/iredis/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|   9 | [py-staticmaps](https://github.com/flopp/py-staticmaps)                    | [2020-09-20](https://github.com/flopp/py-staticmaps/pull/7)                   | [2021-03-24](https://github.com/flopp/py-staticmaps/pull/17)                  | [2](https://github.com/flopp/py-staticmaps/pulls?q=is%3Apr+author%3Ayihong0618)                 |
|  10 | [gitlab-skyline](https://github.com/felixgomez/gitlab-skyline)             | [2021-03-02](https://github.com/felixgomez/gitlab-skyline/pull/6)             | [2021-03-02](https://github.com/felixgomez/gitlab-skyline/pull/6)             | [1](https://github.com/felixgomez/gitlab-skyline/pulls?q=is%3Apr+author%3Ayihong0618)           |
|  11 | [python-garminconnect](https://github.com/cyberjunky/python-garminconnect) | [2021-02-26](https://github.com/cyberjunky/python-garminconnect/pull/43)      | [2021-02-26](https://github.com/cyberjunky/python-garminconnect/pull/43)      | [1](https://github.com/cyberjunky/python-garminconnect/pulls?q=is%3Apr+author%3Ayihong0618)     |
|  12 | [hub-mirror-action](https://github.com/Yikun/hub-mirror-action)            | [2021-04-09](https://github.com/Yikun/hub-mirror-action/pull/101)             | [2021-04-09](https://github.com/Yikun/hub-mirror-action/pull/101)             | [1](https://github.com/Yikun/hub-mirror-action/pulls?q=is%3Apr+author%3Ayihong0618)             |
|  13 | [awesome-cn-cafe-web](https://github.com/antfu/awesome-cn-cafe-web)        | [2020-08-18](https://github.com/antfu/awesome-cn-cafe-web/pull/5)             | [2020-08-18](https://github.com/antfu/awesome-cn-cafe-web/pull/5)             | [1](https://github.com/antfu/awesome-cn-cafe-web/pulls?q=is%3Apr+author%3Ayihong0618)           |
|  14 | [build-your-own-vue](https://github.com/jackiewillen/build-your-own-vue)   | [2020-01-16](https://github.com/jackiewillen/build-your-own-vue/pull/1)       | [2020-01-16](https://github.com/jackiewillen/build-your-own-vue/pull/1)       | [1](https://github.com/jackiewillen/build-your-own-vue/pulls?q=is%3Apr+author%3Ayihong0618)     |
|  15 | [highlight](https://github.com/wenyan-lang/highlight)                      | [2020-09-08](https://github.com/wenyan-lang/highlight/pull/4)                 | [2020-09-08](https://github.com/wenyan-lang/highlight/pull/4)                 | [1](https://github.com/wenyan-lang/highlight/pulls?q=is%3Apr+author%3Ayihong0618)               |
|  16 | [help-to-be-helped](https://github.com/xiaolai/help-to-be-helped)          | [2020-02-04](https://github.com/xiaolai/help-to-be-helped/pull/4)             | [2020-02-04](https://github.com/xiaolai/help-to-be-helped/pull/4)             | [1](https://github.com/xiaolai/help-to-be-helped/pulls?q=is%3Apr+author%3Ayihong0618)           |
|  17 | [GadioVideo](https://github.com/rabbitism/GadioVideo)                      | [2019-09-25](https://github.com/rabbitism/GadioVideo/pull/16)                 | [2019-09-25](https://github.com/rabbitism/GadioVideo/pull/16)                 | [1](https://github.com/rabbitism/GadioVideo/pulls?q=is%3Apr+author%3Ayihong0618)                |
|  18 | [TopList](https://github.com/tophubs/TopList)                              | [2019-08-19](https://github.com/tophubs/TopList/pull/13)                      | [2019-08-19](https://github.com/tophubs/TopList/pull/13)                      | [1](https://github.com/tophubs/TopList/pulls?q=is%3Apr+author%3Ayihong0618)                     |
|  19 | [olo](https://github.com/yetone/olo)                                       | [2021-04-12](https://github.com/yetone/olo/pull/91)                           | [2021-04-12](https://github.com/yetone/olo/pull/91)                           | [1](https://github.com/yetone/olo/pulls?q=is%3Apr+author%3Ayihong0618)                          |
|  20 | [LearnJapan](https://github.com/wizicer/LearnJapan)                        | [2020-03-31](https://github.com/wizicer/LearnJapan/pull/2)                    | [2020-03-31](https://github.com/wizicer/LearnJapan/pull/2)                    | [1](https://github.com/wizicer/LearnJapan/pulls?q=is%3Apr+author%3Ayihong0618)                  |
|  21 | [xrkffgg](https://github.com/xrkffgg/xrkffgg)                              | [2021-03-18](https://github.com/xrkffgg/xrkffgg/pull/3)                       | [2021-03-18](https://github.com/xrkffgg/xrkffgg/pull/3)                       | [1](https://github.com/xrkffgg/xrkffgg/pulls?q=is%3Apr+author%3Ayihong0618)                     |
| sum |                                                                            |                                                                               |                                                                               |                                                                                              55 |

## The repos I stared (random 10)
| ID |                                           REPO                                           | STAREDDATE |     LAUGUAGE     | LATESTUPDATE |
|----|------------------------------------------------------------------------------------------|------------|------------------|--------------|
|  1 | [100-Days-Of-ML-Code](https://github.com/MLEveryday/100-Days-Of-ML-Code)                 | 2018-09-02 | Jupyter Notebook | 2021-04-15   |
|  2 | [shaonianche.github.io](https://github.com/shaonianche/shaonianche.github.io)            | 2020-09-02 | Vue              | 2021-04-13   |
|  3 | [madewithml](https://github.com/GokuMohandas/madewithml)                                 | 2020-04-27 | Jupyter Notebook | 2021-04-14   |
|  4 | [mos-chinadns](https://github.com/IrineSistiana/mos-chinadns)                            | 2020-06-29 | Go               | 2021-04-10   |
|  5 | [shadowsocks-source-analysis](https://github.com/harveyqing/shadowsocks-source-analysis) | 2018-09-25 | md               | 2020-12-08   |
|  6 | [Synonyms](https://github.com/chatopera/Synonyms)                                        | 2018-01-16 | Python           | 2021-04-14   |
|  7 | [run_coursemap](https://github.com/teddokano/run_coursemap)                              | 2021-03-23 | Python           | 2021-03-24   |
|  8 | [go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs)                      | 2018-12-26 | Go               | 2021-03-25   |
|  9 | [kurumi](https://github.com/Hanaasagi/kurumi)                                            | 2021-03-04 | Rust             | 2021-03-05   |
| 10 | [vim-for-server](https://github.com/wklken/vim-for-server)                               | 2019-01-17 | Vim script       | 2021-04-09   |

<!--END_SECTION:my_github-->