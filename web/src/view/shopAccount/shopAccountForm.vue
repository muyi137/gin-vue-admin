<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="平台:" prop="mall_id">
          <el-select v-model="formData.mall_id" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in mallNameOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="用户id:" prop="user_id">
          <el-input v-model.number="formData.user_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Account'
}
</script>

<script setup>
import {
  createAccount,
  updateAccount,
  findAccount
} from '@/api/shopAccount'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const mallNameOptions = ref([])
const formData = ref({
            username: '',
            password: '',
            mall_id: undefined,
            user_id: 0,
        })
// 验证规则
const rule = reactive({
               username : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               password : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mall_id : [{
                   required: true,
                   message: '',
                   trigger: ['select','blur'],
               }]
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAccount({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reaccount
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    mallNameOptions.value = await getDictFunc('mallName')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createAccount(formData.value)
               break
             case 'update':
               res = await updateAccount(formData.value)
               break
             default:
               res = await createAccount(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
