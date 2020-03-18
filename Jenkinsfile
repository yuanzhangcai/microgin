// a monolithic jenkinsfile
// deal with payload sent by gitlab
// resolve gitlab action type among push, merge request, tag
// referenced: https://github.com/jenkinsci/gitlab-plugin
pipeline {
    agent any
    tools {
        go 'Go 1.14'
    }

    environment {
        PACK_FOLDER = "./microgin_pack"
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    try {
                        switch(gitlabActionType) {
                        // case "MERGE":
                        //     echo "repo: ${gitlabSourceRepoHomepage} user: ${gitlabUserName} email:${gitlabUserEmail} action: ${gitlabActionType} source ${gitlabSourceBranch} target ${gitlabTargetBranch}"
                        //     // sh "git checkout ${gitlabSourceBranch}"
                        //     sh "git checkout ${gitlabTargetBranch}"
                        //     sh "git merge origin/${gitlabSourceBranch} -m 'merge'"
                        //     break
                        case "PUSH":
                            // echo "repo: ${gitlabSourceRepoHomepage} user: ${gitlabUserName} email:${gitlabUserEmail} action: ${gitlabActionType} before: ${gitlabBefore} after: ${gitlabAfter}"
                            sh "git checkout ${gitlabAfter}"
                            break
                        case "TAG_PUSH":
                            // echo "repo: ${gitlabSourceRepoHomepage} user: ${gitlabUserName} email:${gitlabUserEmail} action: ${gitlabActionType} before: ${gitlabBefore} after: ${gitlabAfter}"
                            sh "git checkout ${gitlabAfter}"
                            break
                        default:
                            echo gitlabActionType
                        }
                    } catch (Exception ex){
                        echo "push by hand"
                    }
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    sh 'rm -rf ./microgin.tar.gz'
                    sh 'rm -rf ${PACK_FOLDER}'
                    sh 'mkdir ${PACK_FOLDER}'
                    sh 'cp -rf ./etc ${PACK_FOLDER}/'
                    sh 'cp -rf ./script ${PACK_FOLDER}/'
                    if (gitlabActionType == 'TAG_PUSH') {
                        sh 'rm -rf ${PACK_FOLDER}/etc/prod.toml'
                        sh 'GOPROXY=https://goproxy.io; GOSUMDB=off; make prod;'
                    } else {
                        sh 'GOPROXY=https://goproxy.io; GOSUMDB=off; make test;'
                    }
                    sh 'mv microgin ${PACK_FOLDER}/'
                    sh 'tar -czf ./microgin.tar.gz ${PACK_FOLDER}/*'
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    if (gitlabActionType == 'TAG_PUSH') {
                        echo 'TAG_PUSH'
                        NEWEST_TAG = sh(returnStdout: true, script: 'git describe --abbrev=0 --tags').trim()
                        NOWTIME = sh(returnStdout: true, script: "date '+%Y-%m-%d-%H-%M-%S'").trim()
                        FILE_NAME = "microgin_${NEWEST_TAG}_${NOWTIME}.tar.gz"
                        sh "mv microgin.tar.gz ${FILE_NAME}"
                        sh "scp ${FILE_NAME} admin@127.0.0.1:/data/microgin/"
                        sh "mv ${FILE_NAME} ../publish/microgin/"
                        echo "最新包名:${FILE_NAME}"
                    } else {
                        echo 'MERGE'
                        sh 'ssh root@127.0.0.1 "mkdir -p /data/microgin"'
                        sh 'scp ./microgin.tar.gz root@127.0.0.1:/data/microgin/'
                        sh 'ssh root@127.0.0.1 "cd /data/microgin/; rm -rf ${PACK_FOLDER}; tar -xzf microgin.tar.gz;"'
                        sh 'ssh root@127.0.0.1 "cd /data/microgin/; cp -rf ${PACK_FOLDER}/* microgin"'
                        sh 'JENKINS_NODE_COOKIE=dontKillMe; ssh root@127.0.0.1 "/data/microgin/script/restart.sh"'
                    }
                }
            }
        }

        stage('Clean') {
            steps {
                cleanWs()
            }
        }
    }
    post {
        failure {
            dingTalk accessToken: '', jenkinsUrl: 'http://127.0.0.1:10086/',
                    message: "部署失败。", notifyPeople: 'Jenkins'
        }
        // success {
        //     dingTalk accessToken: '', jenkinsUrl: '',
        //             message: "部署成功。", notifyPeople: 'Jenkins'
        // }
    }
}

