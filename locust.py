from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def admin_post01(self):
       self.client.post("/getUser",data={'uuid':'16'})

    @task
    def admin_post02(self):
       self.client.post("/increase",data={'uuid':'16','gifcode':'J5IC0IZ2'})