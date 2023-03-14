module open-cluster-management.io/ocm

go 1.20

replace (
    open-cluster-management.io/registration => ./staging/src/open-cluster-management.io/registration
    open-cluster-management.io/work => ./staging/src/open-cluster-management.io/work
    open-cluster-management.io/placement => ./staging/src/open-cluster-management.io/placement
    open-cluster-management.io/registration-operator => ./staging/src/open-cluster-management.io/registration-operator
)