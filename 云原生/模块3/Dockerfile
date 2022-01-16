FROM centos

ENV APP_RUN_DIR=/opt/app/
RUN mkdir -p ${APP_RUN_DIR}
WORKDIR ${APP_RUN_DIR}
ADD ./k8s_week02  ./
RUN chmod a+x  k8s_week02


CMD ["./k8s_week02"]
