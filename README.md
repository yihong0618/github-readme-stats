# Use

Add a comment to your `README.md` like this:

```md
<!--START_SECTION:my_github-->
<!--END_SECTION:my_github-->
```

- Update your GitHub Actions Workflow permissions

✅ Read and write permissions

> GitHub Menu: Repository's Settings -> Actions -> General -> Workflow permissions

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
          # we need set GITHUB_TOKEN
          GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          # if you also want to send TELE
          TELEGRAM_TOKEN: ${{ secrets.TELE_TOKEN }}
          TELEGRAM_CHAT_ID: ${{ secrets.TELE_CHAT_ID }}
          STARRED_NUM: ${{ env.STARRED_NUM }}

      - name: Push README
        run: |
          git config --local user.email "${{ env.GITHUB_EMAIL }}"
          git config --local user.name "${{ env.GITHUB_NAME }}"
          git commit -a -m 'docs: update readme.md' || echo "nothing to commit"
          git push || echo "nothing to push"
```



# [My example](https://github.com/yihong0618/2021).

## My GitHub Status
<img align="middle" src="https://github-readme-stats-1.yihong0618.vercel.app/api?username=yihong0618&show_icons=true&&&hide_title=true" />

<!--START_SECTION:my_github-->
## The repos I created
| ID  |                                          REPO                                          |   START    |   UPDATE   |     LANGUAGE     | STARS |
|-----|----------------------------------------------------------------------------------------|------------|------------|------------------|-------|
|   1 | [bilingual_book_maker](https://github.com/yihong0618/bilingual_book_maker)             | 2023-03-02 | 2024-12-01 | Python           |  7730 |
|   2 | [xiaogpt](https://github.com/yihong0618/xiaogpt)                                       | 2023-02-16 | 2024-12-01 | Python           |  6295 |
|   3 | [running_page](https://github.com/yihong0618/running_page)                             | 2020-09-17 | 2024-12-01 | Python           |  3627 |
|   4 | [Kindle_download_helper](https://github.com/yihong0618/Kindle_download_helper)         | 2022-06-06 | 2024-11-29 | Python           |  2640 |
|   5 | [GitHubPoster](https://github.com/yihong0618/GitHubPoster)                             | 2021-04-21 | 2024-12-02 | Python           |  1786 |
|   6 | [gitblog](https://github.com/yihong0618/gitblog)                                       | 2019-07-18 | 2024-12-01 | Python           |  1440 |
|   7 | [tg_bot_collections](https://github.com/yihong0618/tg_bot_collections)                 | 2023-12-11 | 2024-11-25 | Python           |   565 |
|   8 | [epubhv](https://github.com/yihong0618/epubhv)                                         | 2023-09-04 | 2024-11-28 | Python           |   470 |
|   9 | [tg_bing_dalle](https://github.com/yihong0618/tg_bing_dalle)                           | 2023-10-04 | 2024-11-25 | Python           |   415 |
|  10 | [MiService](https://github.com/yihong0618/MiService)                                   | 2022-11-04 | 2024-11-28 | Python           |   367 |
|  11 | [twint](https://github.com/yihong0618/twint)                                           | 2021-05-08 | 2024-11-25 | Python           |   325 |
|  12 | [blue](https://github.com/yihong0618/blue)                                             | 2022-10-20 | 2024-11-25 | Python           |   315 |
|  13 | [iBeats](https://github.com/yihong0618/iBeats)                                         | 2021-06-11 | 2024-11-30 | Python           |   293 |
|  14 | [SunoSongsCreator](https://github.com/yihong0618/SunoSongsCreator)                     | 2024-03-23 | 2024-11-09 | Python           |   288 |
|  15 | [2021](https://github.com/yihong0618/2021)                                             | 2020-12-21 | 2024-11-18 | Python           |   273 |
|  16 | [2022](https://github.com/yihong0618/2022)                                             | 2022-01-01 | 2024-08-26 | Python           |   269 |
|  17 | [ChatTTS](https://github.com/yihong0618/ChatTTS)                                       | 2024-05-29 | 2024-11-27 | Jupyter Notebook |   237 |
|  18 | [2023](https://github.com/yihong0618/2023)                                             | 2023-01-01 | 2024-11-09 | Python           |   204 |
|  19 | [shanbay_remember](https://github.com/yihong0618/shanbay_remember)                     | 2020-12-02 | 2024-11-26 | JavaScript       |   191 |
|  20 | [rcode](https://github.com/yihong0618/rcode)                                           | 2022-05-04 | 2024-09-15 | Python           |   189 |
|  21 | [iWhat](https://github.com/yihong0618/iWhat)                                           | 2023-03-08 | 2024-09-04 | Python           |   188 |
|  22 | [duolingo_remember](https://github.com/yihong0618/duolingo_remember)                   | 2021-01-18 | 2024-11-05 | Python           |   166 |
|  23 | [2020](https://github.com/yihong0618/2020)                                             | 2020-01-10 | 2024-10-24 | C                |   136 |
|  24 | [github-readme-stats](https://github.com/yihong0618/github-readme-stats)               | 2020-12-24 | 2024-11-29 | Go               |   135 |
|  25 | [vscode-gcores](https://github.com/yihong0618/vscode-gcores)                           | 2020-01-04 | 2024-11-27 | TypeScript       |   126 |
|  26 | [dalian-IT](https://github.com/yihong0618/dalian-IT)                                   | 2021-04-07 | 2024-11-22 | md               |   116 |
|  27 | [gaycore](https://github.com/yihong0618/gaycore)                                       | 2019-02-18 | 2024-11-12 | Python           |   115 |
|  28 | [klingCreator](https://github.com/yihong0618/klingCreator)                             | 2024-07-12 | 2024-11-18 | Python           |   115 |
|  29 | [IdeoImageCreator](https://github.com/yihong0618/IdeoImageCreator)                     | 2024-03-02 | 2024-11-23 | Python           |    73 |
|  30 | [2024](https://github.com/yihong0618/2024)                                             | 2023-12-29 | 2024-12-01 | Python           |    49 |
|  31 | [yihong0618](https://github.com/yihong0618/yihong0618)                                 | 2020-07-16 | 2024-11-08 | md               |    44 |
|  32 | [kai_xin_ci_chang](https://github.com/yihong0618/kai_xin_ci_chang)                     | 2022-06-15 | 2024-07-11 | Python           |    33 |
|  33 | [BingImageCreator](https://github.com/yihong0618/BingImageCreator)                     | 2023-08-17 | 2024-10-05 | Python           |    33 |
|  34 | [gcores_calendar](https://github.com/yihong0618/gcores_calendar)                       | 2020-08-24 | 2024-12-02 | JavaScript       |    31 |
|  35 | [nbnhhsh-cli](https://github.com/yihong0618/nbnhhsh-cli)                               | 2021-07-09 | 2023-09-05 | Python           |    31 |
|  36 | [pengdu_helper](https://github.com/yihong0618/pengdu_helper)                           | 2021-09-09 | 2024-06-20 | Go               |    29 |
|  37 | [my_kindle_stats](https://github.com/yihong0618/my_kindle_stats)                       | 2021-11-18 | 2023-07-26 | Python           |    24 |
|  38 | [Runtastic](https://github.com/yihong0618/Runtastic)                                   | 2020-07-24 | 2024-10-11 | Python           |    23 |
|  39 | [LumaDreamCreator](https://github.com/yihong0618/LumaDreamCreator)                     | 2024-06-14 | 2024-08-12 | Python           |    22 |
|  40 | [running_skyline](https://github.com/yihong0618/running_skyline)                       | 2021-03-02 | 2024-10-24 | Python           |    21 |
|  41 | [blog](https://github.com/yihong0618/blog)                                             | 2020-06-22 | 2023-01-28 | JavaScript       |    14 |
|  42 | [github-readme-stats-server](https://github.com/yihong0618/github-readme-stats-server) | 2021-12-08 | 2024-07-16 | HTML             |    14 |
|  43 | [Python365](https://github.com/yihong0618/Python365)                                   | 2019-09-05 | 2024-07-16 | Python           |    11 |
|  44 | [run](https://github.com/yihong0618/run)                                               | 2021-08-16 | 2024-12-02 | Python           |    10 |
|  45 | [run_route_show](https://github.com/yihong0618/run_route_show)                         | 2024-11-16 | 2024-12-01 | Python           |     8 |
|  46 | [ask-greenplum](https://github.com/yihong0618/ask-greenplum)                           | 2023-04-10 | 2024-10-21 | Python           |     7 |
|  47 | [pg_geohash](https://github.com/yihong0618/pg_geohash)                                 | 2024-11-05 | 2024-11-08 | Rust             |     5 |
|  48 | [resume](https://github.com/yihong0618/resume)                                         | 2023-10-17 | 2024-11-10 | HTML             |     4 |
|  49 | [pg_polyline](https://github.com/yihong0618/pg_polyline)                               | 2024-10-15 | 2024-10-30 | Rust             |     4 |
|  50 | [remote_mp3_duration](https://github.com/yihong0618/remote_mp3_duration)               | 2024-01-20 | 2024-07-06 | Python           |     3 |
|  51 | [github_upstream_script](https://github.com/yihong0618/github_upstream_script)         | 2021-05-08 | 2022-03-08 | Python           |     2 |
|  52 | [edocteel001](https://github.com/yihong0618/edocteel001)                               | 2019-11-12 | 2022-06-24 | JavaScript       |     1 |
|  53 | [til](https://github.com/yihong0618/til)                                               | 2024-04-24 | 2024-04-28 | md               |     0 |
|  54 | [gaycore-server](https://github.com/yihong0618/gaycore-server)                         | 2019-02-18 | 2020-11-02 | Go               |     0 |
| sum |                                                                                        |            |            |                  | 29512 |

## The repos I contributed to
| ID  |                                         REPO                                         |                                   FIRSTDATE                                    |                                    LASTDATE                                    |  LANGUAGE  |                                              PRCOUNT                                               |
|-----|--------------------------------------------------------------------------------------|--------------------------------------------------------------------------------|--------------------------------------------------------------------------------|------------|----------------------------------------------------------------------------------------------------|
|   1 | [GreenplumPython-archive](https://github.com/greenplum-db/GreenplumPython-archive)   | [2022-03-30](https://github.com/greenplum-db/GreenplumPython-archive/pull/35)  | [2023-07-17](https://github.com/greenplum-db/GreenplumPython-archive/pull/206) | Python     | [25](https://github.com/greenplum-db/GreenplumPython-archive/pulls?q=is%3Apr+author%3Ayihong0618)  |
|   2 | [dify](https://github.com/langgenius/dify)                                           | [2024-11-14](https://github.com/langgenius/dify/pull/10707)                    | [2024-12-01](https://github.com/langgenius/dify/pull/11239)                    | TypeScript | [23](https://github.com/langgenius/dify/pulls?q=is%3Apr+author%3Ayihong0618)                       |
|   3 | [GpxTrackPoster](https://github.com/flopp/GpxTrackPoster)                            | [2019-08-06](https://github.com/flopp/GpxTrackPoster/pull/39)                  | [2021-03-20](https://github.com/flopp/GpxTrackPoster/pull/87)                  | Python     | [12](https://github.com/flopp/GpxTrackPoster/pulls?q=is%3Apr+author%3Ayihong0618)                  |
|   4 | [leetcode-cli](https://github.com/leetcode-tools/leetcode-cli)                       | [2019-11-29](https://github.com/leetcode-tools/leetcode-cli/pull/31)           | [2020-08-21](https://github.com/leetcode-tools/leetcode-cli/pull/49)           | JavaScript | [9](https://github.com/leetcode-tools/leetcode-cli/pulls?q=is%3Apr+author%3Ayihong0618)            |
|   5 | [DB-GPT](https://github.com/eosphoros-ai/DB-GPT)                                     | [2023-05-24](https://github.com/eosphoros-ai/DB-GPT/pull/91)                   | [2023-12-19](https://github.com/eosphoros-ai/DB-GPT/pull/953)                  | Python     | [8](https://github.com/eosphoros-ai/DB-GPT/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|   6 | [uftrace](https://github.com/namhyung/uftrace)                                       | [2023-05-10](https://github.com/namhyung/uftrace/pull/1697)                    | [2024-02-18](https://github.com/namhyung/uftrace/pull/1891)                    | C          | [8](https://github.com/namhyung/uftrace/pulls?q=is%3Apr+author%3Ayihong0618)                       |
|   7 | [autocut](https://github.com/mli/autocut)                                            | [2022-11-17](https://github.com/mli/autocut/pull/43)                           | [2022-11-21](https://github.com/mli/autocut/pull/52)                           | Python     | [7](https://github.com/mli/autocut/pulls?q=is%3Apr+author%3Ayihong0618)                            |
|   8 | [vscode-leetcode](https://github.com/LeetCode-OpenSource/vscode-leetcode)            | [2019-12-03](https://github.com/LeetCode-OpenSource/vscode-leetcode/pull/487)  | [2020-07-22](https://github.com/LeetCode-OpenSource/vscode-leetcode/pull/602)  | TypeScript | [6](https://github.com/LeetCode-OpenSource/vscode-leetcode/pulls?q=is%3Apr+author%3Ayihong0618)    |
|   9 | [risingwave](https://github.com/risingwavelabs/risingwave)                           | [2024-09-25](https://github.com/risingwavelabs/risingwave/pull/18710)          | [2024-10-10](https://github.com/risingwavelabs/risingwave/pull/18845)          | Rust       | [6](https://github.com/risingwavelabs/risingwave/pulls?q=is%3Apr+author%3Ayihong0618)              |
|  10 | [nrc-exporter](https://github.com/yasoob/nrc-exporter)                               | [2020-07-05](https://github.com/yasoob/nrc-exporter/pull/2)                    | [2020-10-07](https://github.com/yasoob/nrc-exporter/pull/11)                   | Python     | [5](https://github.com/yasoob/nrc-exporter/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|  11 | [taichi](https://github.com/taichi-dev/taichi)                                       | [2021-09-23](https://github.com/taichi-dev/taichi/pull/2979)                   | [2021-10-23](https://github.com/taichi-dev/taichi/pull/3256)                   | C++        | [5](https://github.com/taichi-dev/taichi/pulls?q=is%3Apr+author%3Ayihong0618)                      |
|  12 | [Bard-API](https://github.com/dsdanielpark/Bard-API)                                 | [2023-07-19](https://github.com/dsdanielpark/Bard-API/pull/125)                | [2023-07-27](https://github.com/dsdanielpark/Bard-API/pull/142)                | Python     | [5](https://github.com/dsdanielpark/Bard-API/pulls?q=is%3Apr+author%3Ayihong0618)                  |
|  13 | [BingImageCreator](https://github.com/acheong08/BingImageCreator)                    | [2023-04-03](https://github.com/acheong08/BingImageCreator/pull/11)            | [2023-04-07](https://github.com/acheong08/BingImageCreator/pull/17)            | Python     | [4](https://github.com/acheong08/BingImageCreator/pulls?q=is%3Apr+author%3Ayihong0618)             |
|  14 | [awesome-cn-cafe](https://github.com/ElaWorkshop/awesome-cn-cafe)                    | [2020-08-04](https://github.com/ElaWorkshop/awesome-cn-cafe/pull/167)          | [2020-08-10](https://github.com/ElaWorkshop/awesome-cn-cafe/pull/170)          | JavaScript | [3](https://github.com/ElaWorkshop/awesome-cn-cafe/pulls?q=is%3Apr+author%3Ayihong0618)            |
|  15 | [gpdb](https://github.com/beeender/gpdb)                                             | [2022-03-23](https://github.com/beeender/gpdb/pull/1)                          | [2022-05-13](https://github.com/beeender/gpdb/pull/3)                          | C          | [3](https://github.com/beeender/gpdb/pulls?q=is%3Apr+author%3Ayihong0618)                          |
|  16 | [kb](https://github.com/gnebbia/kb)                                                  | [2020-09-21](https://github.com/gnebbia/kb/pull/13)                            | [2020-09-23](https://github.com/gnebbia/kb/pull/28)                            | Python     | [3](https://github.com/gnebbia/kb/pulls?q=is%3Apr+author%3Ayihong0618)                             |
|  17 | [pg-lock-tracer](https://github.com/jnidzwetzki/pg-lock-tracer)                      | [2023-01-28](https://github.com/jnidzwetzki/pg-lock-tracer/pull/26)            | [2023-06-13](https://github.com/jnidzwetzki/pg-lock-tracer/pull/79)            | Python     | [3](https://github.com/jnidzwetzki/pg-lock-tracer/pulls?q=is%3Apr+author%3Ayihong0618)             |
|  18 | [ecapture](https://github.com/gojue/ecapture)                                        | [2022-03-29](https://github.com/gojue/ecapture/pull/15)                        | [2022-05-02](https://github.com/gojue/ecapture/pull/51)                        | C          | [3](https://github.com/gojue/ecapture/pulls?q=is%3Apr+author%3Ayihong0618)                         |
|  19 | [stravalib](https://github.com/stravalib/stravalib)                                  | [2021-08-18](https://github.com/stravalib/stravalib/pull/218)                  | [2022-11-24](https://github.com/stravalib/stravalib/pull/272)                  | Python     | [3](https://github.com/stravalib/stravalib/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|  20 | [python-garminconnect](https://github.com/cyberjunky/python-garminconnect)           | [2021-02-26](https://github.com/cyberjunky/python-garminconnect/pull/43)       | [2021-05-25](https://github.com/cyberjunky/python-garminconnect/pull/49)       | Python     | [2](https://github.com/cyberjunky/python-garminconnect/pulls?q=is%3Apr+author%3Ayihong0618)        |
|  21 | [pgtracer](https://github.com/Aiven-Open/pgtracer)                                   | [2022-11-01](https://github.com/Aiven-Open/pgtracer/pull/27)                   | [2023-06-13](https://github.com/Aiven-Open/pgtracer/pull/43)                   | Python     | [2](https://github.com/Aiven-Open/pgtracer/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|  22 | [nebula-python](https://github.com/vesoft-inc/nebula-python)                         | [2021-05-19](https://github.com/vesoft-inc/nebula-python/pull/106)             | [2021-05-20](https://github.com/vesoft-inc/nebula-python/pull/108)             | Python     | [2](https://github.com/vesoft-inc/nebula-python/pulls?q=is%3Apr+author%3Ayihong0618)               |
|  23 | [paradedb](https://github.com/paradedb/paradedb)                                     | [2023-12-06](https://github.com/paradedb/paradedb/pull/620)                    | [2024-02-19](https://github.com/paradedb/paradedb/pull/875)                    | Rust       | [2](https://github.com/paradedb/paradedb/pulls?q=is%3Apr+author%3Ayihong0618)                      |
|  24 | [pg_badapple](https://github.com/higuoxing/pg_badapple)                              | [2023-09-18](https://github.com/higuoxing/pg_badapple/pull/2)                  | [2023-09-18](https://github.com/higuoxing/pg_badapple/pull/2)                  | C          | [2](https://github.com/higuoxing/pg_badapple/pulls?q=is%3Apr+author%3Ayihong0618)                  |
|  25 | [databend](https://github.com/databendlabs/databend)                                 | [2021-12-29](https://github.com/databendlabs/databend/pull/3690)               | [2021-12-30](https://github.com/databendlabs/databend/pull/3701)               | Rust       | [2](https://github.com/databendlabs/databend/pulls?q=is%3Apr+author%3Ayihong0618)                  |
|  26 | [garth](https://github.com/matin/garth)                                              | [2023-09-26](https://github.com/matin/garth/pull/5)                            | [2023-09-27](https://github.com/matin/garth/pull/7)                            | Python     | [2](https://github.com/matin/garth/pulls?q=is%3Apr+author%3Ayihong0618)                            |
|  27 | [py-staticmaps](https://github.com/flopp/py-staticmaps)                              | [2020-09-20](https://github.com/flopp/py-staticmaps/pull/7)                    | [2021-03-24](https://github.com/flopp/py-staticmaps/pull/17)                   | Python     | [2](https://github.com/flopp/py-staticmaps/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|  28 | [iredis](https://github.com/laixintao/iredis)                                        | [2019-12-30](https://github.com/laixintao/iredis/pull/184)                     | [2020-09-16](https://github.com/laixintao/iredis/pull/360)                     | Python     | [2](https://github.com/laixintao/iredis/pulls?q=is%3Apr+author%3Ayihong0618)                       |
|  29 | [activities](https://github.com/flopp/activities)                                    | [2020-07-09](https://github.com/flopp/activities/pull/41)                      | [2020-07-14](https://github.com/flopp/activities/pull/44)                      | JavaScript | [2](https://github.com/flopp/activities/pulls?q=is%3Apr+author%3Ayihong0618)                       |
|  30 | [opendal](https://github.com/apache/opendal)                                         | [2023-01-12](https://github.com/apache/opendal/pull/1179)                      | [2023-02-03](https://github.com/apache/opendal/pull/1269)                      | Rust       | [2](https://github.com/apache/opendal/pulls?q=is%3Apr+author%3Ayihong0618)                         |
|  31 | [strava-datasource](https://github.com/grafana/strava-datasource)                    | [2021-04-13](https://github.com/grafana/strava-datasource/pull/34)             | [2021-05-13](https://github.com/grafana/strava-datasource/pull/39)             | TypeScript | [2](https://github.com/grafana/strava-datasource/pulls?q=is%3Apr+author%3Ayihong0618)              |
|  32 | [viztracer](https://github.com/gaogaotiantian/viztracer)                             | [2024-03-18](https://github.com/gaogaotiantian/viztracer/pull/413)             | [2024-07-24](https://github.com/gaogaotiantian/viztracer/pull/454)             | Python     | [2](https://github.com/gaogaotiantian/viztracer/pulls?q=is%3Apr+author%3Ayihong0618)               |
|  33 | [hub-mirror-action](https://github.com/Yikun/hub-mirror-action)                      | [2021-04-09](https://github.com/Yikun/hub-mirror-action/pull/101)              | [2021-04-19](https://github.com/Yikun/hub-mirror-action/pull/106)              | Python     | [2](https://github.com/Yikun/hub-mirror-action/pulls?q=is%3Apr+author%3Ayihong0618)                |
|  34 | [bs-core](https://github.com/zu1k/bs-core)                                           | [2022-11-30](https://github.com/zu1k/bs-core/pull/25)                          | [2022-11-30](https://github.com/zu1k/bs-core/pull/25)                          | TypeScript | [1](https://github.com/zu1k/bs-core/pulls?q=is%3Apr+author%3Ayihong0618)                           |
|  35 | [MangaLineExtraction_PyTorch](https://github.com/ljsabc/MangaLineExtraction_PyTorch) | [2021-09-22](https://github.com/ljsabc/MangaLineExtraction_PyTorch/pull/3)     | [2021-09-22](https://github.com/ljsabc/MangaLineExtraction_PyTorch/pull/3)     | Python     | [1](https://github.com/ljsabc/MangaLineExtraction_PyTorch/pulls?q=is%3Apr+author%3Ayihong0618)     |
|  36 | [pyWhat](https://github.com/bee-san/pyWhat)                                          | [2021-06-18](https://github.com/bee-san/pyWhat/pull/89)                        | [2021-06-18](https://github.com/bee-san/pyWhat/pull/89)                        | Python     | [1](https://github.com/bee-san/pyWhat/pulls?q=is%3Apr+author%3Ayihong0618)                         |
|  37 | [tetos](https://github.com/frostming/tetos)                                          | [2024-04-18](https://github.com/frostming/tetos/pull/2)                        | [2024-04-18](https://github.com/frostming/tetos/pull/2)                        | Python     | [1](https://github.com/frostming/tetos/pulls?q=is%3Apr+author%3Ayihong0618)                        |
|  38 | [Striker](https://github.com/s0md3v/Striker)                                         | [2019-06-20](https://github.com/s0md3v/Striker/pull/64)                        | [2019-06-20](https://github.com/s0md3v/Striker/pull/64)                        | Python     | [1](https://github.com/s0md3v/Striker/pulls?q=is%3Apr+author%3Ayihong0618)                         |
|  39 | [awesome-database-learning](https://github.com/pingcap/awesome-database-learning)    | [2021-05-11](https://github.com/pingcap/awesome-database-learning/pull/37)     | [2021-05-11](https://github.com/pingcap/awesome-database-learning/pull/37)     | md         | [1](https://github.com/pingcap/awesome-database-learning/pulls?q=is%3Apr+author%3Ayihong0618)      |
|  40 | [fugashi](https://github.com/polm/fugashi)                                           | [2023-12-28](https://github.com/polm/fugashi/pull/87)                          | [2023-12-28](https://github.com/polm/fugashi/pull/87)                          | C++        | [1](https://github.com/polm/fugashi/pulls?q=is%3Apr+author%3Ayihong0618)                           |
|  41 | [cloudberry-site](https://github.com/apache/cloudberry-site)                         | [2024-11-04](https://github.com/apache/cloudberry-site/pull/169)               | [2024-11-04](https://github.com/apache/cloudberry-site/pull/169)               | TypeScript | [1](https://github.com/apache/cloudberry-site/pulls?q=is%3Apr+author%3Ayihong0618)                 |
|  42 | [pgvectorscale](https://github.com/timescale/pgvectorscale)                          | [2024-06-13](https://github.com/timescale/pgvectorscale/pull/94)               | [2024-06-13](https://github.com/timescale/pgvectorscale/pull/94)               | Rust       | [1](https://github.com/timescale/pgvectorscale/pulls?q=is%3Apr+author%3Ayihong0618)                |
|  43 | [pdbattach](https://github.com/jschwinger233/pdbattach)                              | [2022-03-07](https://github.com/jschwinger233/pdbattach/pull/1)                | [2022-03-07](https://github.com/jschwinger233/pdbattach/pull/1)                | Python     | [1](https://github.com/jschwinger233/pdbattach/pulls?q=is%3Apr+author%3Ayihong0618)                |
|  44 | [build-your-own-vue](https://github.com/jackiewillen/build-your-own-vue)             | [2020-01-16](https://github.com/jackiewillen/build-your-own-vue/pull/1)        | [2020-01-16](https://github.com/jackiewillen/build-your-own-vue/pull/1)        | JavaScript | [1](https://github.com/jackiewillen/build-your-own-vue/pulls?q=is%3Apr+author%3Ayihong0618)        |
|  45 | [generative-ai-python](https://github.com/google-gemini/generative-ai-python)        | [2023-12-14](https://github.com/google-gemini/generative-ai-python/pull/115)   | [2023-12-14](https://github.com/google-gemini/generative-ai-python/pull/115)   | Python     | [1](https://github.com/google-gemini/generative-ai-python/pulls?q=is%3Apr+author%3Ayihong0618)     |
|  46 | [help-to-be-helped](https://github.com/xiaolai/help-to-be-helped)                    | [2020-02-04](https://github.com/xiaolai/help-to-be-helped/pull/4)              | [2020-02-04](https://github.com/xiaolai/help-to-be-helped/pull/4)              | md         | [1](https://github.com/xiaolai/help-to-be-helped/pulls?q=is%3Apr+author%3Ayihong0618)              |
|  47 | [tzf-rs](https://github.com/ringsaturn/tzf-rs)                                       | [2023-01-25](https://github.com/ringsaturn/tzf-rs/pull/36)                     | [2023-01-25](https://github.com/ringsaturn/tzf-rs/pull/36)                     | Rust       | [1](https://github.com/ringsaturn/tzf-rs/pulls?q=is%3Apr+author%3Ayihong0618)                      |
|  48 | [tokei-pie](https://github.com/laixintao/tokei-pie)                                  | [2021-11-19](https://github.com/laixintao/tokei-pie/pull/2)                    | [2021-11-19](https://github.com/laixintao/tokei-pie/pull/2)                    | Python     | [1](https://github.com/laixintao/tokei-pie/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|  49 | [derive](https://github.com/erik/derive)                                             | [2019-09-28](https://github.com/erik/derive/pull/28)                           | [2019-09-28](https://github.com/erik/derive/pull/28)                           | JavaScript | [1](https://github.com/erik/derive/pulls?q=is%3Apr+author%3Ayihong0618)                            |
|  50 | [awesome-cn-cafe-web](https://github.com/antfu/awesome-cn-cafe-web)                  | [2020-08-18](https://github.com/antfu/awesome-cn-cafe-web/pull/5)              | [2020-08-18](https://github.com/antfu/awesome-cn-cafe-web/pull/5)              | TypeScript | [1](https://github.com/antfu/awesome-cn-cafe-web/pulls?q=is%3Apr+author%3Ayihong0618)              |
|  51 | [TopList](https://github.com/tophubs/TopList)                                        | [2019-08-19](https://github.com/tophubs/TopList/pull/13)                       | [2019-08-19](https://github.com/tophubs/TopList/pull/13)                       | Go         | [1](https://github.com/tophubs/TopList/pulls?q=is%3Apr+author%3Ayihong0618)                        |
|  52 | [postgresml](https://github.com/postgresml/postgresml)                               | [2023-07-07](https://github.com/postgresml/postgresml/pull/810)                | [2023-07-07](https://github.com/postgresml/postgresml/pull/810)                | Rust       | [1](https://github.com/postgresml/postgresml/pulls?q=is%3Apr+author%3Ayihong0618)                  |
|  53 | [pgvector](https://github.com/pgvector/pgvector)                                     | [2023-06-08](https://github.com/pgvector/pgvector/pull/148)                    | [2023-06-08](https://github.com/pgvector/pgvector/pull/148)                    | C          | [1](https://github.com/pgvector/pgvector/pulls?q=is%3Apr+author%3Ayihong0618)                      |
|  54 | [pgrx](https://github.com/pgcentralfoundation/pgrx)                                  | [2023-06-29](https://github.com/pgcentralfoundation/pgrx/pull/1181)            | [2023-06-29](https://github.com/pgcentralfoundation/pgrx/pull/1181)            | Rust       | [1](https://github.com/pgcentralfoundation/pgrx/pulls?q=is%3Apr+author%3Ayihong0618)               |
|  55 | [xrkffgg](https://github.com/xrkffgg/xrkffgg)                                        | [2021-03-18](https://github.com/xrkffgg/xrkffgg/pull/3)                        | [2021-03-18](https://github.com/xrkffgg/xrkffgg/pull/3)                        | JavaScript | [1](https://github.com/xrkffgg/xrkffgg/pulls?q=is%3Apr+author%3Ayihong0618)                        |
|  56 | [nebula](https://github.com/vesoft-inc/nebula)                                       | [2021-05-17](https://github.com/vesoft-inc/nebula/pull/2476)                   | [2021-05-17](https://github.com/vesoft-inc/nebula/pull/2476)                   | C++        | [1](https://github.com/vesoft-inc/nebula/pulls?q=is%3Apr+author%3Ayihong0618)                      |
|  57 | [GadioVideo](https://github.com/rabbitism/GadioVideo)                                | [2019-09-25](https://github.com/rabbitism/GadioVideo/pull/16)                  | [2019-09-25](https://github.com/rabbitism/GadioVideo/pull/16)                  | Python     | [1](https://github.com/rabbitism/GadioVideo/pulls?q=is%3Apr+author%3Ayihong0618)                   |
|  58 | [notion-avatar](https://github.com/Mayandev/notion-avatar)                           | [2021-09-28](https://github.com/Mayandev/notion-avatar/pull/1)                 | [2021-09-28](https://github.com/Mayandev/notion-avatar/pull/1)                 | TypeScript | [1](https://github.com/Mayandev/notion-avatar/pulls?q=is%3Apr+author%3Ayihong0618)                 |
|  59 | [pgvecto.rs](https://github.com/tensorchord/pgvecto.rs)                              | [2023-06-28](https://github.com/tensorchord/pgvecto.rs/pull/33)                | [2023-06-28](https://github.com/tensorchord/pgvecto.rs/pull/33)                | Rust       | [1](https://github.com/tensorchord/pgvecto.rs/pulls?q=is%3Apr+author%3Ayihong0618)                 |
|  60 | [DingdangD1-PoC](https://github.com/LynMoe/DingdangD1-PoC)                           | [2022-08-17](https://github.com/LynMoe/DingdangD1-PoC/pull/2)                  | [2022-08-17](https://github.com/LynMoe/DingdangD1-PoC/pull/2)                  | Python     | [1](https://github.com/LynMoe/DingdangD1-PoC/pulls?q=is%3Apr+author%3Ayihong0618)                  |
|  61 | [juicesync](https://github.com/juicedata/juicesync)                                  | [2021-12-29](https://github.com/juicedata/juicesync/pull/119)                  | [2021-12-29](https://github.com/juicedata/juicesync/pull/119)                  | Go         | [1](https://github.com/juicedata/juicesync/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|  62 | [tzfpy](https://github.com/ringsaturn/tzfpy)                                         | [2023-01-25](https://github.com/ringsaturn/tzfpy/pull/7)                       | [2023-01-25](https://github.com/ringsaturn/tzfpy/pull/7)                       | Python     | [1](https://github.com/ringsaturn/tzfpy/pulls?q=is%3Apr+author%3Ayihong0618)                       |
|  63 | [MiService](https://github.com/Yonsm/MiService)                                      | [2022-11-04](https://github.com/Yonsm/MiService/pull/10)                       | [2022-11-04](https://github.com/Yonsm/MiService/pull/10)                       | Python     | [1](https://github.com/Yonsm/MiService/pulls?q=is%3Apr+author%3Ayihong0618)                        |
|  64 | [cloudberry](https://github.com/apache/cloudberry)                                   | [2024-11-08](https://github.com/apache/cloudberry/pull/707)                    | [2024-11-08](https://github.com/apache/cloudberry/pull/707)                    | C          | [1](https://github.com/apache/cloudberry/pulls?q=is%3Apr+author%3Ayihong0618)                      |
|  65 | [PgNodeGraph](https://github.com/charliettxx/PgNodeGraph)                            | [2023-05-15](https://github.com/charliettxx/PgNodeGraph/pull/1)                | [2023-05-15](https://github.com/charliettxx/PgNodeGraph/pull/1)                | md         | [1](https://github.com/charliettxx/PgNodeGraph/pulls?q=is%3Apr+author%3Ayihong0618)                |
|  66 | [laixintao](https://github.com/laixintao/laixintao)                                  | [2023-08-23](https://github.com/laixintao/laixintao/pull/1)                    | [2023-08-23](https://github.com/laixintao/laixintao/pull/1)                    | Python     | [1](https://github.com/laixintao/laixintao/pulls?q=is%3Apr+author%3Ayihong0618)                    |
|  67 | [github-repos-stats](https://github.com/DaoCloud-OpenSource/github-repos-stats)      | [2022-04-06](https://github.com/DaoCloud-OpenSource/github-repos-stats/pull/4) | [2022-04-06](https://github.com/DaoCloud-OpenSource/github-repos-stats/pull/4) | Go         | [1](https://github.com/DaoCloud-OpenSource/github-repos-stats/pulls?q=is%3Apr+author%3Ayihong0618) |
|  68 | [olo](https://github.com/yetone/olo)                                                 | [2021-04-12](https://github.com/yetone/olo/pull/91)                            | [2021-04-12](https://github.com/yetone/olo/pull/91)                            | Python     | [1](https://github.com/yetone/olo/pulls?q=is%3Apr+author%3Ayihong0618)                             |
|  69 | [highlight](https://github.com/wenyan-lang/highlight)                                | [2020-09-08](https://github.com/wenyan-lang/highlight/pull/4)                  | [2020-09-08](https://github.com/wenyan-lang/highlight/pull/4)                  | JavaScript | [1](https://github.com/wenyan-lang/highlight/pulls?q=is%3Apr+author%3Ayihong0618)                  |
|  70 | [openmodelz](https://github.com/tensorchord/openmodelz)                              | [2023-08-16](https://github.com/tensorchord/openmodelz/pull/160)               | [2023-08-16](https://github.com/tensorchord/openmodelz/pull/160)               | Go         | [1](https://github.com/tensorchord/openmodelz/pulls?q=is%3Apr+author%3Ayihong0618)                 |
|  71 | [DreamMachineAPI](https://github.com/danaigc/DreamMachineAPI)                        | [2024-06-14](https://github.com/danaigc/DreamMachineAPI/pull/3)                | [2024-06-14](https://github.com/danaigc/DreamMachineAPI/pull/3)                | Python     | [1](https://github.com/danaigc/DreamMachineAPI/pulls?q=is%3Apr+author%3Ayihong0618)                |
|  72 | [pymc](https://github.com/pymc-devs/pymc)                                            | [2023-07-12](https://github.com/pymc-devs/pymc/pull/6823)                      | [2023-07-12](https://github.com/pymc-devs/pymc/pull/6823)                      | Python     | [1](https://github.com/pymc-devs/pymc/pulls?q=is%3Apr+author%3Ayihong0618)                         |
|  73 | [rust-xgboost](https://github.com/postgresml/rust-xgboost)                           | [2023-07-07](https://github.com/postgresml/rust-xgboost/pull/6)                | [2023-07-07](https://github.com/postgresml/rust-xgboost/pull/6)                | Rust       | [1](https://github.com/postgresml/rust-xgboost/pulls?q=is%3Apr+author%3Ayihong0618)                |
|  74 | [LearnJapan](https://github.com/wizicer/LearnJapan)                                  | [2020-03-31](https://github.com/wizicer/LearnJapan/pull/2)                     | [2020-03-31](https://github.com/wizicer/LearnJapan/pull/2)                     | TypeScript | [1](https://github.com/wizicer/LearnJapan/pulls?q=is%3Apr+author%3Ayihong0618)                     |
|  75 | [cloudberry-gpbackup](https://github.com/apache/cloudberry-gpbackup)                 | [2024-11-04](https://github.com/apache/cloudberry-gpbackup/pull/9)             | [2024-11-04](https://github.com/apache/cloudberry-gpbackup/pull/9)             | Go         | [1](https://github.com/apache/cloudberry-gpbackup/pulls?q=is%3Apr+author%3Ayihong0618)             |
|  76 | [incubator-devlake](https://github.com/apache/incubator-devlake)                     | [2021-11-23](https://github.com/apache/incubator-devlake/pull/846)             | [2021-11-23](https://github.com/apache/incubator-devlake/pull/846)             | Go         | [1](https://github.com/apache/incubator-devlake/pulls?q=is%3Apr+author%3Ayihong0618)               |
|  77 | [UsePythonProcessDataFaster](https://github.com/linw1995/UsePythonProcessDataFaster) | [2021-09-01](https://github.com/linw1995/UsePythonProcessDataFaster/pull/1)    | [2021-09-01](https://github.com/linw1995/UsePythonProcessDataFaster/pull/1)    | md         | [1](https://github.com/linw1995/UsePythonProcessDataFaster/pulls?q=is%3Apr+author%3Ayihong0618)    |
|  78 | [gitlab-skyline](https://github.com/felixgomez/gitlab-skyline)                       | [2021-03-02](https://github.com/felixgomez/gitlab-skyline/pull/6)              | [2021-03-02](https://github.com/felixgomez/gitlab-skyline/pull/6)              | Python     | [1](https://github.com/felixgomez/gitlab-skyline/pulls?q=is%3Apr+author%3Ayihong0618)              |
|  79 | [community](https://github.com/eosphoros-ai/community)                               | [2023-08-03](https://github.com/eosphoros-ai/community/pull/1)                 | [2023-08-03](https://github.com/eosphoros-ai/community/pull/1)                 | md         | [1](https://github.com/eosphoros-ai/community/pulls?q=is%3Apr+author%3Ayihong0618)                 |
| sum |                                                                                      |                                                                                |                                                                                |            |                                                                                                215 |
<!--END_SECTION:my_github-->
