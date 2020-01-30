<template>
  <div class="app-container">
    <el-row>
      <el-button type="primary" size="mini" style="float: right; margin-bottom: 10px"
                 @click="newUserDialogVisible = true">
        添加用户
      </el-button>
    </el-row>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID">
        <template slot-scope="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="账号">
        <template slot-scope="scope">
          {{ scope.row.account }}
        </template>
      </el-table-column>
      <el-table-column label="注册时间" align="center" width="170px">
        <template slot-scope="scope">
          {{ parseTime(scope.row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-row>
            <el-button type="primary" size="mini" @click="onPasswordUpdate(scope.row.id)">
              修改密码
            </el-button>
          </el-row>
          <br/>
          <el-row>
            <el-button type="danger" size="mini" @click="onDelete(scope.row.id)">删除</el-button>
          </el-row>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination style="float: right; margin-top: 10px"
                   background
                   layout="prev, pager, next"
                   :total=total
                   :page-sizes=perPage
    >
    </el-pagination>

    <!--    修改密码-->
    <el-dialog title="修改密码" :visible.sync="updatePasswordDialogVisible">
      <el-form :model="form" label-position="right" label-width="60px">
        <el-form-item label="新密码">
          <el-input v-model="form.password" autocomplete="off"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="updatePasswordDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="onPasswordSubmit()">确 定</el-button>
      </div>
    </el-dialog>

    <!--    添加用户-->
    <el-dialog title="添加用户" :visible.sync="newUserDialogVisible">
      <el-form :model="form" label-position="right" label-width="60px">
        <el-form-item label="账号">
          <el-input v-model="form.account" autocomplete="off"/>
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="form.password" autocomplete="off"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="newUserDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="onNewUserSubmit()">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
  import { getList, update, register, destroy } from '@/api/user'
  import { parseTime } from '@/utils/index'

  export default {
    data() {
      return {
        list: null,
        listLoading: true,
        perPage: [],
        total: 0,
        updatePasswordDialogVisible: false,
        newUserDialogVisible: false,
        form: {},
        userID: 0
      }
    },
    created() {
      this.fetchData()
    },
    methods: {
      fetchData() {
        this.listLoading = true
        getList().then(response => {
          this.list = response.data
          this.total = response.pagination.total
          this.perPage = [response.pagination.per_page]
          this.listLoading = false
        })
      },
      parseTime(data) {
        return parseTime(data)
      },
      onDelete(id) {
        this.$confirm('确认删除此用户?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {

          destroy(id).then(response => {
            this.$message({
              message: response.message,
              type: 'success'
            })

            this.fetchData()
          })

        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
      },
      onPasswordUpdate(id) {
        this.userID = id
        this.updatePasswordDialogVisible = true
      },
      onPasswordSubmit() {
        update(this.userID, {
          password: this.form.password
        }).then(response => {
          this.$message({
            message: response.message,
            type: 'success'
          })
          this.updatePasswordDialogVisible = false
        })
      },
      onNewUserSubmit() {
        register(this.form).then(response => {
          this.$message({
            message: response.message,
            type: 'success'
          })

          this.fetchData()

          this.newUserDialogVisible = false
        })
      }
    }
  }
</script>
