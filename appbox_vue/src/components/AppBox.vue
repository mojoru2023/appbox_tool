<template>
<div class="container" id="vueApp">
    <div class="row mx-auto w-75" style="margin-top: 20px">
        <div class="col-12 text-center">
            <h3>App Box 🦉</h3>
        </div>
    </div>
    <div class="row mx-auto w-75" style="height: 15px"></div>
    <div class="row mx-auto w-75">
        <div class="col-6">
            <div class="btn-group">
    <button class="btn btn-outline-info btn-sm" @click="create()">新增</button>
    <div class="modal fade" tabindex="-1" ref="modal1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h4 class="modal-title">新增应用信息</h4>
            <button
              type="button"
              class="btn-close"
              data-bs-dismiss="modal"
              aria-label="Close"
            ></button>
          </div>
          <div class="modal-body">
                    <div class="row">
                            <div class="col-3">应用名称：</div>
                            <div class="col-9">
                                <input class="form-control" placeholder="应用名称" v-model="newRow.app_name">
                            </div>
                        </div>

                                            <div class="row">
                            <div class="col-3">应用类型：</div>
                            <div class="col-9">
                                <input class="form-control" placeholder="应用类型" v-model="newRow.app_type">
                            </div>
                        </div>

                        <div class="row">
                            <div class="col-3">域名：</div>
                            <div class="col-9">
                                <input class="form-control" placeholder="域名" v-model="newRow.domain">
                            </div>
                        </div>

                        <div class="row">
                            <div class="col-3">技术栈：</div>
                            <div class="col-9">
                                <input class="form-control" placeholder="技术栈" v-model="newRow.stack">
                            </div>
                        </div>
                                            <div class="row">
                            <div class="col-3">前端端口：</div>
                            <div class="col-9">
                                <input class="form-control" placeholder="前端端口" v-model="newRow.font_port">
                            </div>
                        </div>
                                            <div class="row">
                            <div class="col-3">后端端口：</div>
                            <div class="col-9">
                                <input class="form-control" placeholder="后端端口" v-model="newRow.backend_port">
                            </div>
                        </div>
                                            <div class="row">
                            <div class="col-3">其他：</div>
                            <div class="col-9">
                                <input class="form-control" placeholder="其他" v-model="newRow.others">
                            </div>
                        </div>


                    </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-outline-primary" data-dismiss="modal" @click="addRow">确认</button>
            <button
              type="button"
              class="btn btn-secondary"
              data-bs-dismiss="modal"
            >
              取消
            </button>

          </div>
        </div>
      </div>
    </div>
            </div>
            <button type="button" class="btn btn-outline-warning btn-sm" @click="delRows">删除</button>
        </div>
        <div class="col-6">
            <div class="input-group">
                <input type="text" class="form-control input-group-sm" placeholder="输入应用id进行搜索">
                <span class="input-group-btn">
                        <button class="btn btn-default" type="button"><i class="fa fa-search"></i></button>
                    </span>
            </div>
        </div>
    </div>
    <div class="row mx-auto w-75" style="height: 15px"></div>
    <div class="row mx-auto w-75">
        <div class="col-22">
            <table class="table table-hover table-success">
                <thead class="thead-default">
                <tr>
                    <th><input type="checkbox"></th>
                   <th>序号</th>
                    <th>应用名称</th>
                    <th>应用类型</th>
                    <th>域名</th>
                    <th>技术栈</th>
                    <th>前端端口</th>
                    <th>后端端口</th>


                </tr>
                </thead>
                <tbody>

                <tr v-for=" (app,index)  in apps" :key="(app, index)">

                    <td><input type="checkbox" :value="index" v-model="checkedRows"></td>
                    <td>{{app.id}}</td>
                    <td>{{app.app_name}}</td>
                    <td>{{app.app_type}}</td>
                    <td>{{app.domain}}</td>
                    <td>{{app.stack}}</td>
                    <td>{{app.fontport}}</td>
                    <td>{{app.backendport}}</td>


                </tr>
                </tbody>
            </table>
        </div>
    </div>

</div>
</template>

<script>
import axios from "axios";
import { Modal } from 'bootstrap';
export default {
    name: "AppBox",
    props:['id', 'app_name', 'app_type', 'domain', 'stack', 'fontport', 'backendport'],
    provide() {
        return {
            reload:this.reload
        }
    },

    data() {
        return {
            checkAll: false,
            checkedRows: [],
            apps: "",
            newRow: {},
                  myModal1: null,
        };
    },
    inject: ['reload'],
    methods: {
            create: function () {
            this.isEidt = false;
            this.myModal1.show();
        },
        addRow: function () {
            this.apps.push(this.newRow);
            let data = new FormData();
            data.append("app_name", this.newRow.app_name)
            axios.post("/v1/appbox/api", {
                "app_name": this.newRow.app_name,
                "app_type": this.newRow.app_type,
                "domain": this.newRow.domain,
                "stack": this.newRow.stack,
                "fontport": this.newRow.font_port,
                "backendport": this.newRow.backend_port,
            }).then(
                (res) => {
                    console.info(res);
                    this.reload();
                    this.$notify({
                        group: 'foo',
                        type: 'success',
                        text: '添加成功'
                    });
                }).catch(
                    (error) => {
                        console.error(error)
                    })
            console.log(this.newRow);
            this.newRow = {};
            location.reload(); // 重新加载
        },
        delRows: function () {
            if (this.checkedRows.length <= 0) {
                alert("您未选择需要删除的数据");
                return false;
            }
            if (!confirm("您确定要删除选择的数据吗？")) {
                return false;
            }
            var remove_id = String(this.apps[this.checkedRows[0]]["id"]);
            console.log(remove_id);
            axios.delete('/v1/appbox/api' + '/' + remove_id).then(() => {
                this.reload();
                this.$notify({
                    group: 'foo',
                    type: 'info',
                    text: '删除成功'
                });
                this.reload()
            }).catch((error) => {
                console.log(error)
            })
            location.reload(); // 删除后重新加载
        },
        getAllDT: function () {
            axios.get('/v1/appbox/api').then((res) => {
                this.apps = res.data;
            }).catch((error) => {
                console.log(error)
            });
        }
    },
    mounted() {
        this.getAllDT();
        this.myModal1 = new Modal(this.$refs.modal1, { keyboard: true });

    }
}
</script>
