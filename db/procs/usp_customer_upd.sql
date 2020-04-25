USE [GoKube]
GO

/****** Object:  StoredProcedure [Customer].[usp_customer_upd]    Script Date: 4/23/2020 1:38:04 PM ******/
IF EXISTS(SELECT *
FROM sys.procedures
WHERE name = 'usp_customer_upd')
BEGIN
    DROP PROCEDURE [Customer].[usp_customer_upd]
END
GO

/****** Object:  StoredProcedure [Customer].[usp_customer_upd]    Script Date: 4/23/2020 1:38:04 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


/****************************************************************************

DATE        VERSION       AUTHOR                REFERENCE

2020-04-23  1.0           Naveen Prabhaker

Initial Version

*****************************************************************************/

CREATE PROCEDURE [Customer].[usp_customer_upd]
    (
    @CustomerID int,
    @NAME nvarchar(100) NULL,
    @InternalCustomerID nvarchar(100) NULL,
    @StartDate datetime NULL,
    @EndDate datetime NULL,
    @Rowversion rowversion NULL,
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
        FROM customer.customer
        WHERE CustomerID=@CustomerID AND Rowversion=@Rowversion) AND @UnitTestBypass=0
  BEGIN
        RAISERROR('This record has been updated by another user.', 16,1);
        RETURN;
    END

    IF EXISTS (SELECT 1
        FROM customer.customer
        WHERE CustomerID <> @CustomerID AND Name = @NAME ) AND @UnitTestBypass=0
  BEGIN
        RAISERROR('Cannot update Customer Name. Name already exist.', 16,1);
        RETURN;
    END

    IF EXISTS (SELECT 1
    FROM customer.Customer
    WHERE Name = @NAME AND InternalCustomerID = @InternalCustomerID)
  BEGIN

        SELECT
            @NAME = ISNULL(@NAME, Name),
            @InternalCustomerID = ISNULL(@InternalCustomerID, InternalCustomerID),
            @StartDate = ISNULL(@StartDate, StartDate),
            @EndDate = ISNULL(@EndDate, EndDate)
        FROM
            Customer.Customer
        WHERE CustomerID = @CustomerID

        UPDATE [Customer].[Customer]
	SET
		[Name] = @NAME,
		InternalCustomerID = @InternalCustomerID,
		StartDate = @StartDate,
		EndDate = @EndDate
	WHERE
		CustomerID = @CustomerID
            AND (Rowversion=@Rowversion OR (Rowversion!=@Rowversion AND @UnitTestBypass=1))


    END
  ELSE
  BEGIN

        RAISERROR('Record Does not exists',16,1);

    END

END

GO


