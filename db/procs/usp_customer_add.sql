USE [GoKube]
GO

/****** Object:  StoredProcedure [Customer].[usp_customer_add]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_customer_add')
BEGIN
    DROP PROCEDURE [Customer].[usp_customer_add]
END
GO

/****** Object:  StoredProcedure [Customer].[usp_customer_add]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-04-23  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [Customer].[usp_customer_add]
    (
    @NAME nvarchar(100),
    @InternalCustomerID nvarchar(100),
    @StartDate datetime,
    @EndDate datetime,
    @UnitTestBypass bit=0
)
AS
BEGIN

    SET NOCOUNT ON;

    IF @@TRANCOUNT = 0
    BEGIN
        RAISERROR('Procedure must be called within a transaction', 16,1);
        RETURN;
    END

    IF NOT EXISTS (SELECT 1
    FROM customer.Customer
    WHERE Name = @NAME AND InternalCustomerID = @InternalCustomerID)
    BEGIN
        DECLARE @CustomerID INT
        SELECT @CustomerID = ISNULL(MAX(CustomerID), 0 ) + 1
        FROM Customer.Customer

        INSERT INTO [Customer].[Customer]
            ([CustomerID]
            ,[Name]
            ,[InternalCustomerID]
            ,[StartDate]
            ,[EndDate]
            ,[InsertedDate])
        VALUES
            (
                @CustomerID
		   , @NAME
           , @InternalCustomerID
           , @StartDate
           , @EndDate
           , GETUTCDATE())

    END
    ELSE
    BEGIN
        RAISERROR('Record already exists',16,1);
    END

END

GO


