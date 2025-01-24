// 1. Find High-Spending Users
db.users.aggregate([
    {
      $lookup: {
        from: "orders",
        localField: "userId",
        foreignField: "userId",
        as: "userOrders"
      }
    },
    {
      $unwind: "$userOrders"
    },
    {
      $group: {
        _id: "$userId",
        totalSpent: { $sum: "$userOrders.totalAmount" },
        name: { $first: "$name" },
        email: { $first: "$email" }
      }
    },
    {
      $match: {
        totalSpent: { $gt: 500 }
      }
    },
    {
      $project: {
        _id: 0,
        userId: "$_id",
        name: 1,
        email: 1,
        totalSpent: 1
      }
    }
  ]);
  
  // 2. List Popular Products by Average Rating
  db.products.aggregate([
    {
      $unwind: "$ratings"
    },
    {
      $group: {
        _id: "$productId",
        name: { $first: "$name" },
        category: { $first: "$category" },
        avgRating: { $avg: "$ratings.rating" }
      }
    },
    {
      $match: {
        avgRating: { $gte: 4 }
      }
    },
    {
      $project: {
        _id: 0,
        productId: "$_id",
        name: 1,
        category: 1,
        avgRating: 1
      }
    }
  ]);
  
  // 3. Search for Orders in a Specific Time Range
  db.orders.aggregate([
    {
      $match: {
        orderDate: {
          $gte: new ISODate("2024-12-01T00:00:00Z"),
          $lte: new ISODate("2024-12-31T23:59:59Z")
        }
      }
    },
    {
      $lookup: {
        from: "users",
        localField: "userId",
        foreignField: "userId",
        as: "userDetails"
      }
    },
    {
      $unwind: "$userDetails"
    },
    {
      $project: {
        _id: 0,
        orderId: 1,
        orderDate: 1,
        totalAmount: 1,
        status: 1,
        userName: "$userDetails.name"
      }
    }
  ]);
  
  // 4. Update Stock After Order Completion
  db.orders.find({ orderId: "ORD001" }).forEach(order => {
    order.items.forEach(item => {
      db.products.updateOne(
        { productId: item.productId },
        { $inc: { stock: -item.quantity } }
      );
    });
  });
  
  // 5. Find Nearest Warehouse
  db.warehouses.aggregate([
    {
      $geoNear: {
        near: { type: "Point", coordinates: [-74.006, 40.7128] },
        distanceField: "distance",
        maxDistance: 50000,
        query: { products: "P001" },
        spherical: true
      }
    },
    {
      $project: {
        _id: 0,
        warehouseId: 1,
        location: 1,
        distance: 1,
        products: 1
      }
    }
  ]);
  