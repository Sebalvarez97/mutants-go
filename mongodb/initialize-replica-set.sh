kubectl exec mongod-0 bash
root@mongod-0:/# hostname -f
root@mongod-0:/# mongo
> rs.initiate({ _id: "MainRepSet", version: 1,
members: [
 { _id: 0, host: "mongod-0.mongodb-service.mutants.svc.cluster.local:27017" },
 { _id: 1, host: "mongod-1.mongodb-service.mutants.svc.cluster.local:27017" },
 { _id: 2, host: "mongod-2.mongodb-service.mutants.svc.cluster.local:27017" } ]});
MainRepSet:PRIMARY> rs.status();
MainRepSet:PRIMARY> db.getSiblingDB("admin").createUser({
...       user : "admin",
...       pwd  : "admin1234",
...       roles: [ { role: "root", db: "admin" } ]
...  });