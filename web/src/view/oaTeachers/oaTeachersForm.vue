<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="学院:" prop="college">
          <el-select v-model="formData.college" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in collegeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="部门:" prop="department">
          <el-input v-model="formData.department" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:" prop="gender">
          <el-select v-model="formData.gender" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="身份证号:" prop="cardNumber">
          <el-input v-model="formData.cardNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="职务:" prop="position">
          <el-select v-model="formData.position" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in positionOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="学历:" prop="qualifications">
          <el-select v-model="formData.qualifications" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in education backgroundOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="职称:" prop="professional">
          <el-select v-model="formData.professional" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in professional titleOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="来校时间:" prop="hiredate">
          <el-date-picker v-model="formData.hiredate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="岗位:" prop="jobs">
          <el-select v-model="formData.jobs" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in jobsOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'OaTeachers'
}
</script>

<script setup>
import {
  createOaTeachers,
  updateOaTeachers,
  findOaTeachers
} from '@/api/oaTeachers'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const collegeOptions = ref([])
const genderOptions = ref([])
const positionOptions = ref([])
const education backgroundOptions = ref([])
const professional titleOptions = ref([])
const jobsOptions = ref([])
const formData = ref({
            college: undefined,
            department: '',
            name: '',
            gender: undefined,
            cardNumber: '',
            position: undefined,
            qualifications: undefined,
            professional: undefined,
            hiredate: new Date(),
            jobs: undefined,
        })
// 验证规则
const rule = reactive({
               college : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cardNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               position : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               qualifications : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               professional : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               hiredate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               jobs : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOaTeachers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reoaTeachers
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    collegeOptions.value = await getDictFunc('college')
    genderOptions.value = await getDictFunc('gender')
    positionOptions.value = await getDictFunc('position')
    education backgroundOptions.value = await getDictFunc('education background')
    professional titleOptions.value = await getDictFunc('professional title')
    jobsOptions.value = await getDictFunc('jobs')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOaTeachers(formData.value)
               break
             case 'update':
               res = await updateOaTeachers(formData.value)
               break
             default:
               res = await createOaTeachers(formData.value)
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
